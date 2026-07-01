with open('src/main.cpp', 'r') as f:
    content = f.read()

content = content.replace('const static std::wstring friendly_app_name = WSTRING(L"Piano Game " << PianoGameVersionString);', '')
content = content.replace('friendly_app_name', 'WSTRING(L"Piano Game " << PianoGameVersionString)')

with open('src/main.cpp', 'w') as f:
    f.write(content)

with open('src/CompatibleSystem.cpp', 'r') as f:
    content = f.read()

content = content.replace('const static std::wstring friendly_app_name = WSTRING(L"Piano Game " << PianoGameVersionString);\n      const static std::wstring message_box_title = WSTRING(friendly_app_name << L" Error");', 'const static std::wstring message_box_title = WSTRING(L"Piano Game " << PianoGameVersionString << L" Error");')
content = content.replace('const static std::wstring friendly_app_name = WSTRING(L"Piano Game " << PianoGameVersionString);\r\n      const static std::wstring message_box_title = WSTRING(friendly_app_name << L" Error");', 'const static std::wstring message_box_title = WSTRING(L"Piano Game " << PianoGameVersionString << L" Error");')

with open('src/CompatibleSystem.cpp', 'w') as f:
    f.write(content)
