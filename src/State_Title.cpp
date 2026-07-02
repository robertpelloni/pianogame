
// Copyright (c)2007 Nicholas Piegdon
// See license.txt for license information

#include "State_Title.h"
#include "State_Playing.h"

#include "version.h"
#include "CompatibleSystem.h"

#include "MenuLayout.h"
#include "UserSettings.h"
#include "file_selector.h"
#include "Renderer.h"
#include "Textures.h"

#include "libmidi/Midi.h"
#include "libmidi/MidiUtil.h"
#include "libmidi/MidiComm.h"

using namespace std;

const static wstring OutputDeviceKey = L"Last Output Device";
const static wstring OutputKeySpecialDisabled = L"[no output device]";

const static wstring InputDeviceKey = L"Last Input Device";
const static wstring InputKeySpecialDisabled = L"[no input device]";

TitleState::~TitleState()
{
   if (m_output_tile) delete m_output_tile;
   if (m_input_tile) delete m_input_tile;
   if (m_file_tile) delete m_file_tile;
}


void TitleState::RebuildTrackTiles()
{
   m_track_tiles.clear();
   if (!m_state.midi) return;

   int track_count = 0;
   for (size_t i = 0; i < m_state.midi->Tracks().size(); ++i)
   {
      if (m_state.midi->Tracks()[i].Notes().size()) track_count++;
   }

   int tiles_across = (GetStateWidth() / 2) / (TrackTileWidth + Layout::ScreenMarginX);
   tiles_across = std::max(tiles_across, 1);

   int tiles_down = (GetStateHeight() - Layout::ScreenMarginX - Layout::ScreenMarginY * 2) / (TrackTileHeight + Layout::ScreenMarginX);
   tiles_down = std::max(tiles_down, 1);

   m_tiles_per_page = tiles_across * tiles_down;

   m_page_count        = track_count / m_tiles_per_page;
   const int remainder = track_count % m_tiles_per_page;
   if (remainder > 0) m_page_count++;

   if (track_count < tiles_across) tiles_across = track_count;

   int all_tile_widths = tiles_across * TrackTileWidth + (tiles_across-1) * Layout::ScreenMarginX;
   int global_x_offset = GetStateWidth() / 2 + (GetStateWidth() / 2 - all_tile_widths) / 2;

   const static int starting_y = 100;

   int tiles_on_this_line = 0;
   int tiles_on_this_page = 0;
   int current_y = starting_y;
   for (size_t i = 0; i < m_state.midi->Tracks().size(); ++i)
   {
      const MidiTrack &t = m_state.midi->Tracks()[i];
      if (t.Notes().size() == 0) continue;

      int x = global_x_offset + (TrackTileWidth + Layout::ScreenMarginX)*tiles_on_this_line;
      int y = current_y;

      Track::Mode mode = Track::ModePlayedAutomatically;
      if (t.IsPercussion()) mode = Track::ModePlayedButHidden;

      Track::TrackColor color = static_cast<Track::TrackColor>((m_track_tiles.size()) % Track::UserSelectableColorCount);

      if (m_state.track_properties.size() > i)
      {
         color = m_state.track_properties[i].color;
         mode = m_state.track_properties[i].mode;
      }

      TrackTile tile(x, y, i, color, mode);

      m_track_tiles.push_back(tile);

      tiles_on_this_line++;
      tiles_on_this_line %= tiles_across;
      if (!tiles_on_this_line)
      {
         current_y += TrackTileHeight + Layout::ScreenMarginX;
      }

      tiles_on_this_page++;
      tiles_on_this_page %= m_tiles_per_page;
      if (!tiles_on_this_page)
      {
         current_y = starting_y;
         tiles_on_this_line = 0;
      }
   }
}

std::vector<Track::Properties> TitleState::BuildTrackProperties() const
{
   std::vector<Track::Properties> props;
   if (!m_state.midi) return props;

   for (size_t i = 0; i < m_state.midi->Tracks().size(); ++i)
   {
      props.push_back(Track::Properties());
   }

   for (std::vector<TrackTile>::const_iterator i = m_track_tiles.begin(); i != m_track_tiles.end(); ++i)
   {
      props[i->GetTrackId()].color = i->GetColor();
      props[i->GetTrackId()].mode = i->GetMode();
   }

   return props;
}


void TitleState::Init()
{
   m_back_button = ButtonState(Layout::ScreenMarginX,
      GetStateHeight() - Layout::ScreenMarginY/2 - Layout::ButtonHeight/2,
      Layout::ButtonWidth, Layout::ButtonHeight);

   m_continue_button = ButtonState(GetStateWidth() - Layout::ScreenMarginX - Layout::ButtonWidth,
      GetStateHeight() - Layout::ScreenMarginY/2 - Layout::ButtonHeight/2,
      Layout::ButtonWidth, Layout::ButtonHeight);

   wstring last_output_device = UserSetting::Get(OutputDeviceKey, L"");
   wstring last_input_device = UserSetting::Get(InputDeviceKey, L"");

   // midi_out could be in one of three states right now:
   //    1. We just started and were passed a null MidiCommOut pointer
   // Or, we're returning from track selection either with:
   //    2. a null MidiCommOut because the user didn't want any output, or
   //    3. a valid MidiCommOut we constructed previously.
   if (!m_state.midi_out)
   {
      // Try to find the previously used device
      MidiCommDescriptionList devices = MidiCommOut::GetDeviceList();
      for (size_t i = 0; i < devices.size(); ++i)
      {
         if (devices[i].name == last_output_device)
         {
            m_state.midi_out = new MidiCommOut(devices[i].id);
         }
      }

      // Next, if we couldn't find a previously used device,
      // use the first one
      if (last_output_device != OutputKeySpecialDisabled && !m_state.midi_out && devices.size() > 0)
      {
         m_state.midi_out = new MidiCommOut(devices[0].id);
      }
   }

   if (!m_state.midi_in)
   {
      // Try to find the previously used device
      MidiCommDescriptionList devices = MidiCommIn::GetDeviceList();
      for (size_t i = 0; i < devices.size(); ++i)
      {
         if (devices[i].name == last_input_device)
         {
            try
            {
               m_state.midi_in = new MidiCommIn(devices[i].id);
            }
            catch (MidiErrorCode)
            {
               m_state.midi_in = 0;
            }
         }
      }

      // If we couldn't find the previously used device,
      // disabling by default (i.e. leaving it null) is
      // completely acceptable.
   }

   int output_device_id = -1;
   if (last_output_device == OutputKeySpecialDisabled) output_device_id = -1;

   if (m_state.midi_out)
   {
      output_device_id = m_state.midi_out->GetDeviceDescription().id;
      m_state.midi_out->Reset();
   }

   int input_device_id = -1;
   if (m_state.midi_in)
   {
      input_device_id = m_state.midi_in->GetDeviceDescription().id;
      m_state.midi_in->Reset();
   }

   const bool compress_height = (GetStateHeight() < 750);
   const int initial_y = (compress_height ? 230 : 360);
   const int each_y = (compress_height ? 94 : 100);

   m_file_tile = new StringTile((GetStateWidth() - StringTileWidth) / 2, initial_y + each_y*0, GetTexture(SongBox));
   m_file_tile->SetString(m_state.song_title);

   const MidiCommDescriptionList output_devices = MidiCommOut::GetDeviceList();
   const MidiCommDescriptionList input_devices = MidiCommIn::GetDeviceList();
   m_output_tile = new DeviceTile((GetStateWidth() - DeviceTileWidth) / 2, initial_y + each_y*1, output_device_id, DeviceTileOutput, output_devices, GetTexture(InterfaceButtons), GetTexture(OutputBox));
   m_input_tile = new DeviceTile((GetStateWidth() - DeviceTileWidth) / 2, initial_y + each_y*2, input_device_id, DeviceTileInput, input_devices, GetTexture(InterfaceButtons), GetTexture(InputBox));
}

void TitleState::Update()
{
   MouseInfo mouse = Mouse();

   if (m_skip_next_mouse_up)
   {
      mouse.released.left = false;
      m_skip_next_mouse_up = false;
   }

   m_continue_button.Update(mouse);
   m_back_button.Update(mouse);

   MouseInfo output_mouse(mouse);
   output_mouse.x -= m_output_tile->GetX();
   output_mouse.y -= m_output_tile->GetY();
   m_output_tile->Update(output_mouse);

   MouseInfo input_mouse(mouse);
   input_mouse.x -= m_input_tile->GetX();
   input_mouse.y -= m_input_tile->GetY();
   m_input_tile->Update(input_mouse);

   MouseInfo file_mouse(mouse);
   file_mouse.x -= m_file_tile->GetX();
   file_mouse.y -= m_file_tile->GetY();
   m_file_tile->Update(file_mouse);

   // Check to see for clicks on the file box
   if (m_file_tile->Hit())
   {
      m_skip_next_mouse_up = true;

      if (m_state.midi_out)
      {
         m_state.midi_out->Reset();
         m_output_tile->TurnOffPreview();
      }

      Midi *new_midi = 0;

      std::wstring filename;
      std::wstring file_title;
      FileSelector::RequestMidiFilename(&filename, &file_title);

      if (filename != L"")
      {
         try
         {
            new_midi = new Midi(Midi::ReadFromFile(filename));
         }
         catch (const MidiError &e)
         {
            wstring wrapped_description = WSTRING(L"Problem while loading file: " << file_title << L"\n") + e.GetErrorDescription();
            Compatible::ShowError(wrapped_description);

            new_midi = 0;
         }

         if (new_midi)
         {
            SharedState new_state;
            new_state.midi = new_midi;
            new_state.midi_in = m_state.midi_in;
            new_state.midi_out = m_state.midi_out;
            new_state.song_title = FileSelector::TrimFilename(filename);

            delete m_state.midi;
            m_state = new_state;

            m_file_tile->SetString(m_state.song_title);
            RebuildTrackTiles();
         }
      }
   }

   // Check to see if we need to switch to a newly selected output device
   int output_id = m_output_tile->GetDeviceId();
   if (!m_state.midi_out || output_id != static_cast<int>(m_state.midi_out->GetDeviceDescription().id))
   {
      if (m_state.midi_out) m_state.midi_out->Reset();

      delete m_state.midi_out;
      m_state.midi_out = 0;

      if (output_id >= 0)
      {
         m_state.midi_out = new MidiCommOut(output_id);
         m_state.midi->Reset(0,0);

         UserSetting::Set(OutputDeviceKey, m_state.midi_out->GetDeviceDescription().name);
      }
      else
      {
         UserSetting::Set(OutputDeviceKey, OutputKeySpecialDisabled);
      }
   }

   if (m_state.midi_out)
   {
      if (m_output_tile->HitPreviewButton())
      {
         m_state.midi_out->Reset();

         if (m_output_tile->IsPreviewOn())
         {
            const microseconds_t PreviewLeadIn  = 250000;
            const microseconds_t PreviewLeadOut = 250000;
            m_state.midi->Reset(PreviewLeadIn, PreviewLeadOut);

            PlayDevicePreview(0);
         }
      }
      else
      {
         PlayDevicePreview(static_cast<microseconds_t>(GetDeltaMilliseconds()) * 1000);
      }
   }


   int input_id = m_input_tile->GetDeviceId();
   if (!m_state.midi_in || input_id != static_cast<int>(m_state.midi_in->GetDeviceDescription().id))
   {
      if (m_state.midi_in) m_state.midi_in->Reset();
      m_last_input_note_name = "";

      delete m_state.midi_in;
      m_state.midi_in = 0;

      if (input_id >= 0)
      {
         try
         {
            m_state.midi_in = new MidiCommIn(input_id);
            UserSetting::Set(InputDeviceKey, m_state.midi_in->GetDeviceDescription().name);
         }
         catch (MidiErrorCode)
         {
            m_state.midi_in = 0;
         }
      }
      else
      {
         UserSetting::Set(InputDeviceKey, InputKeySpecialDisabled);
      }
   }

   if (m_state.midi_in && m_input_tile->IsPreviewOn())
   {
      // Read note events to display on screen
      while (m_state.midi_in->KeepReading())
      {
         MidiEvent ev = m_state.midi_in->Read();
         if (ev.Type() == MidiEventType_NoteOff || ev.Type() == MidiEventType_NoteOn)
         {
            string note = MidiEvent::NoteName(ev.NoteNumber());

            if (ev.Type() == MidiEventType_NoteOn && ev.NoteVelocity() > 0)
            {
               m_last_input_note_name = note;
            }
            else
            {
               if (note == m_last_input_note_name) m_last_input_note_name = "";
            }
         }
      }
   }
   else
   {
      m_last_input_note_name = "";
   }


   if (IsKeyPressed(KeyEscape) || m_back_button.hit)
   {
      delete m_state.midi_out;
      m_state.midi_out = 0;

      delete m_state.midi_in;
      m_state.midi_in = 0;

      delete m_state.midi;
      m_state.midi = 0;

      Compatible::GracefulShutdown();
      return;
   }

   if (IsKeyPressed(KeyEnter) || m_continue_button.hit)
   {
      if (!m_state.midi) return;
      if (m_state.midi_out) m_state.midi_out->Reset();
      if (m_state.midi_in) m_state.midi_in->Reset();

      m_state.track_properties = BuildTrackProperties();
      ChangeState(new PlayingState(m_state));
      return;
   }

   if (IsKeyPressed(KeyDown) || IsKeyPressed(KeyRight))
   {
      m_current_page++;
      if (m_current_page == m_page_count) m_current_page = 0;
   }

   if (IsKeyPressed(KeyUp) || IsKeyPressed(KeyLeft))
   {
      m_current_page--;
      if (m_current_page < 0) m_current_page += m_page_count;
   }

   m_tooltip = L"";

   if (m_back_button.hovering) m_tooltip = L"Quit the application entirely.";
   if (m_continue_button.hovering) m_tooltip = L"Save configuration and begin playing the game.";

   if (!m_first_update_after_seek)
   {
      PlayTrackPreview(static_cast<microseconds_t>(GetDeltaMilliseconds()) * 1000);
   }
   m_first_update_after_seek = false;

   if (m_track_tiles.size() > 0)
   {
       size_t start = m_current_page * m_tiles_per_page;
       size_t end = std::min( static_cast<size_t>((m_current_page+1) * m_tiles_per_page), m_track_tiles.size() );
       for (size_t i = start; i < end; ++i)
       {
          TrackTile &t = m_track_tiles[i];

          MouseInfo tm = MouseInfo(mouse);
          tm.x -= t.GetX();
          tm.y -= t.GetY();

          t.Update(tm);

          if (t.ButtonLeft().hovering || t.ButtonRight().hovering)
          {
             switch (t.GetMode())
             {
             case Track::ModeNotPlayed: m_tooltip = L"Track won't be played or shown during the game."; break;
             case Track::ModePlayedAutomatically: m_tooltip = L"Track will be played automatically by the game."; break;
             case Track::ModePlayedButHidden: m_tooltip = L"Track will be played automatically by the game, but also hidden from view."; break;
             case Track::ModeYouPlay: m_tooltip = L"'You Play' means you want to play this track yourself."; break;
             }
          }

          if (t.ButtonPreview().hovering)
          {
             if (t.IsPreviewOn()) m_tooltip = L"Turn track preview off.";
             else m_tooltip = L"Preview how this track sounds.";
          }

          if (t.ButtonColor().hovering) m_tooltip = L"Customize the visual note color for this track's lane.";

          if (t.HitPreviewButton())
          {
             if (m_state.midi_out) m_state.midi_out->Reset();

             if (t.IsPreviewOn())
             {
                for (size_t j = 0; j < m_track_tiles.size(); ++j)
                {
                   if (i == j) continue;
                   m_track_tiles[j].TurnOffPreview();
                }

                const microseconds_t PreviewLeadIn  = 25000;
                const microseconds_t PreviewLeadOut = 25000;

                m_preview_on = true;
                m_preview_track_id = t.GetTrackId();
                m_state.midi->Reset(PreviewLeadIn, PreviewLeadOut);
                PlayTrackPreview(0);

                microseconds_t additional_time = -PreviewLeadIn;
                const MidiTrack &track = m_state.midi->Tracks()[m_preview_track_id];
                for (size_t k = 0; k < track.Events().size(); ++k)
                {
                   const MidiEvent &ev = track.Events()[k];
                   if (ev.Type() == MidiEventType_NoteOn && ev.NoteVelocity() > 0)
                   {
                      additional_time += track.EventUsecs()[k] - m_state.midi->GetDeadAirStartOffsetMicroseconds() - 1;
                      break;
                   }
                }

                PlayTrackPreview(additional_time);
                m_first_update_after_seek = true;
             }
             else
             {
                m_preview_on = false;
             }
          }
       }
   }

   if (m_file_tile->WholeTile().hovering) m_tooltip = L"Open the file browser to load a custom .mid / .midi track.";

   if (m_input_tile->ButtonLeft().hovering) m_tooltip = L"Select the MIDI hardware input (e.g. your piano keyboard) to record strokes.";
   if (m_input_tile->ButtonRight().hovering) m_tooltip = L"Select the MIDI hardware input (e.g. your piano keyboard) to record strokes.";
   if (m_input_tile->ButtonPreview().hovering)
   {
      if (m_input_tile->IsPreviewOn()) m_tooltip = L"End diagnostic mode for this input device.";
      else m_tooltip = L"Start diagnostic mode: press keys on your piano to test input visibility.";
   }

   if (m_output_tile->ButtonLeft().hovering) m_tooltip = L"Select the MIDI hardware output (e.g. software synth) to hear playback.";
   if (m_output_tile->ButtonRight().hovering) m_tooltip = L"Select the MIDI hardware output (e.g. software synth) to hear playback.";
   if (m_output_tile->ButtonPreview().hovering)
   {
      if (m_output_tile->IsPreviewOn()) m_tooltip = L"End diagnostic mode for this output device.";
      else m_tooltip = L"Start diagnostic mode: outputs sample chords to verify sound functionality.";
   }

}

void TitleState::PlayDevicePreview(microseconds_t delta_microseconds)
{
   if (!m_output_tile->IsPreviewOn()) return;
   if (!m_state.midi_out) return;

   MidiEventListWithTrackId evs = m_state.midi->Update(delta_microseconds);

   for (MidiEventListWithTrackId::const_iterator i = evs.begin(); i != evs.end(); ++i)
   {
      m_state.midi_out->Write(i->second);
   }
}

void TitleState::PlayTrackPreview(microseconds_t delta_microseconds)
{
   if (!m_preview_on) return;

   MidiEventListWithTrackId evs = m_state.midi->Update(delta_microseconds);

   for (MidiEventListWithTrackId::const_iterator i = evs.begin(); i != evs.end(); ++i)
   {
      const MidiEvent &ev = i->second;
      if (i->first != m_preview_track_id) continue;

      if (m_state.midi_out) m_state.midi_out->Write(ev);
   }
}

void TitleState::Draw(Renderer &renderer) const
{
   const bool compress_height = (GetStateHeight() < 750);
   const bool compress_width = (GetStateWidth() < 850);

   const static int TitleWidth = 507;
   const static int TitleY = (compress_height ? 20 : 100);

   int left_pane_center_x = GetStateWidth() / 4;
   int left = left_pane_center_x - TitleWidth / 2;

   renderer.SetColor(White);
   renderer.DrawTga(GetTexture(TitleLogo), left, TitleY);

   TextWriter version(Layout::ScreenMarginX, GetStateHeight() - Layout::ScreenMarginY - Layout::SmallFontSize * 2,
      renderer, false, Layout::SmallFontSize);

   std::wstring extra = L"";
#ifdef _DEBUG
   extra = L" debug";
#endif

   version << Text(WSTRING(L"version " << PianoGameVersionString << extra), Gray);

   Layout::DrawHorizontalRule(renderer, GetStateWidth(), GetStateHeight() - Layout::ScreenMarginY);

   // Title state is now the dashboard, so continue means play song.
   Layout::DrawButton(renderer, m_continue_button, GetTexture(ButtonPlaySong));
   Layout::DrawButton(renderer, m_back_button, GetTexture(ButtonExit));

   m_output_tile->Draw(renderer);
   m_input_tile->Draw(renderer);
   m_file_tile->Draw(renderer);

   if (m_input_tile->IsPreviewOn())
   {
      const static int PreviewWidth = 60;
      const static int PreviewHeight = 40;

      const int x = m_input_tile->GetX() + DeviceTileWidth + 12;
      const int y = m_input_tile->GetY() + 38;

      renderer.SetColor(0xFF, 0xFF, 0xFF);
      renderer.DrawQuad(x, y, PreviewWidth, PreviewHeight);

      renderer.SetColor(0,0,0);
      renderer.DrawQuad(x+1, y+1, PreviewWidth-2, PreviewHeight-2);

      TextWriter last_note(x + PreviewWidth/2 - 1, m_input_tile->GetY() + 44, renderer, true, Layout::TitleFontSize);
      Widen<wchar_t> w;
      last_note << w(m_last_input_note_name);
   }

   if (m_track_tiles.size() > 0)
   {
      // Write our page count on the screen
      TextWriter pagination(GetStateWidth()/2 + GetStateWidth()/4, GetStateHeight() - Layout::SmallFontSize - 30, renderer, true, Layout::ButtonFontSize);
      pagination << Text(WSTRING(L"Page " << (m_current_page+1) << L" of " << m_page_count << L" (arrow keys change page)"), Gray);

      Tga *buttons = GetTexture(InterfaceButtons);
      Tga *box = GetTexture(TrackPanel);

      size_t start = m_current_page * m_tiles_per_page;
      size_t end = std::min( static_cast<size_t>((m_current_page+1) * m_tiles_per_page), m_track_tiles.size() );
      for (size_t i = start; i < end; ++i)
      {
         m_track_tiles[i].Draw(renderer, m_state.midi, buttons, box);
      }
   }

   const int tooltip_font_size = (compress_width ? Layout::ButtonFontSize : Layout::TitleFontSize);
   TextWriter tooltip(GetStateWidth() / 2, GetStateHeight() - Layout::ScreenMarginY/2 - tooltip_font_size/2, renderer, true, tooltip_font_size);
   tooltip << m_tooltip;
}
