
// Copyright (c)2007 Nicholas Piegdon
// See license.txt for license information

#include "PianoGameError.h"
#include "string_util.h"

using namespace std;

const std::wstring error_header1 = L"Piano Game detected a";
const std::wstring error_header2 = L" problem and must close:\n\n";
const std::wstring error_footer = L"\n";

std::wstring PianoGameError::GetErrorDescription() const
{
   switch (m_error)
   {
   case Error_StringSpecified:             return m_optional_string;

   case Error_BadPianoType:                return L"Bad piano type specified.";
   case Error_BadGameState:                return L"Internal Error: Piano Game entered bad game state!";

   default:                                return WSTRING(L"Unknown PianoGameError Code (" << m_error << L").");
   }
}
