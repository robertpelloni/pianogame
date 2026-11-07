use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct WindowConfig {
    pub title: String,
    pub width: u32,
    pub height: u32,
    pub resizable: bool,
    pub tray_docked: bool,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DesktopContainer {
    pub window: WindowConfig,
    pub url: String,
}

impl DesktopContainer {
    pub fn new() -> Self {
        DesktopContainer {
            window: WindowConfig {
                title: "Ultimate Agentic Coding Harness".to_string(),
                width: 1280,
                height: 850,
                resizable: true,
                tray_docked: true,
            },
            url: "http://localhost:8080".to_string(),
        }
    }

    pub fn get_launch_info(&self) -> String {
        format!(
            "Desktop Container initialized: {} ({}x{}) docked={}",
            self.window.title, self.window.width, self.window.height, self.window.tray_docked
        )
    }
}
