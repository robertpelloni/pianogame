
// Copyright (c)2007 Nicholas Piegdon
// See license.txt for license information

#ifndef __PIANOGAME_VERSION_H
#define __PIANOGAME_VERSION_H

#include <string>

// See readme.txt for a list of what has changed between versions
#include "string_util.h"

static const std::wstring PianoGameVersionString = L"0.6.9";

#define GET_PIANO_GAME_FRIENDLY_APP_NAME WSTRING(L"Piano Game " << PianoGameVersionString)

#endif
