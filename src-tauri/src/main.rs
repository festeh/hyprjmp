// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use tauri::{CustomMenuItem, Manager, SystemTray, SystemTrayMenu};

fn main() {
    let quit = CustomMenuItem::new("open".to_string(), "Open");
    let tray_menu = SystemTrayMenu::new().add_item(quit);
    let tray = SystemTray::new().with_menu(tray_menu);
    tauri::Builder::default()
        .system_tray(tray)
        .on_system_tray_event(|app, event| match event {
            tauri::SystemTrayEvent::MenuItemClick { .. } => {
                let window = app.get_window("main").unwrap();
                window.hide().unwrap();
                window.show().unwrap();
            }
            _ => {
                println!("Other event");
            }
        })
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
