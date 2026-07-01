# Deployment

## Windows
Build the project using Visual Studio (PianoGame.vcproj). NSIS is used to build the installer executable (`nsis_installer_script.nsi`).

## Mac OS X
Build using `Synthesia.xcodeproj`.

## Automation
- A custom script or Python tool should be used to bump versions universally across `version.h`, `nsis_installer_script.nsi`, `VERSION.md`, and `CHANGELOG.md` to prevent hard-coding misalignments.
