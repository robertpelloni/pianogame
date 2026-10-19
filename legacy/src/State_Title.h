
// Copyright (c)2007 Nicholas Piegdon
// See license.txt for license information

#ifndef __STATE_TITLE_H
#define __STATE_TITLE_H

#include "SharedState.h"
#include "GameState.h"
#include "MenuLayout.h"
#include "libmidi/MidiTypes.h"
#include "DeviceTile.h"
#include "StringTile.h"
#include "TrackTile.h"
#include <vector>

class Midi;
class MidiCommOut;
class Tga;

class TitleState : public GameState
{
public:
   TitleState(const SharedState &state)
      : m_state(state), m_output_tile(0), m_input_tile(0),
        m_file_tile(0), m_skip_next_mouse_up(false),
        m_preview_on(false), m_first_update_after_seek(false), m_preview_track_id(0),
        m_page_count(0), m_current_page(0), m_tiles_per_page(0)
   { }

   ~TitleState();

protected:
   virtual void Init();
   virtual void Update();
   virtual void Draw(Renderer &renderer) const;

private:
   void PlayDevicePreview(microseconds_t delta_microseconds);
   void PlayTrackPreview(microseconds_t additional_time);
   std::vector<Track::Properties> BuildTrackProperties() const;
   void RebuildTrackTiles();

   ButtonState m_continue_button;
   ButtonState m_back_button;

   SharedState m_state;

   std::string m_last_input_note_name;
   std::wstring m_tooltip;

   DeviceTile *m_output_tile;
   DeviceTile *m_input_tile;
   StringTile *m_file_tile;

   std::vector<TrackTile> m_track_tiles;
   int m_page_count;
   int m_current_page;
   int m_tiles_per_page;
   bool m_preview_on;
   bool m_first_update_after_seek;
   size_t m_preview_track_id;

   bool m_skip_next_mouse_up;
};

#endif
