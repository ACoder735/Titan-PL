#include <stdio.h>
#include <stdlib.h>
#include <windows.h>
#include "utils.h"

void print_error(const char* message, int line) {
    HANDLE hConsole = GetStdHandle(STD_OUTPUT_HANDLE);
    
    // 1. Save current color
    CONSOLE_SCREEN_BUFFER_INFO consoleInfo;
    GetConsoleScreenBufferInfo(hConsole, &consoleInfo);
    WORD saved_attributes = consoleInfo.wAttributes;

    // 2. Set Color to Bright Red
    SetConsoleTextAttribute(hConsole, FOREGROUND_RED | FOREGROUND_INTENSITY);

    // 3. Print ENTIRE message in Red
    if (line > 0) {
        printf("[Error] Line %d: %s\n", line, message);
    } else {
        printf("[Error]: %s\n", message);
    }

    // 4. Reset color back to normal
    SetConsoleTextAttribute(hConsole, saved_attributes);
    
    // 5. Exit
    exit(1);
}

void print_warning(const char* message, int line) {
    HANDLE hConsole = GetStdHandle(STD_OUTPUT_HANDLE);
    
    CONSOLE_SCREEN_BUFFER_INFO consoleInfo;
    GetConsoleScreenBufferInfo(hConsole, &consoleInfo);
    WORD saved_attributes = consoleInfo.wAttributes;

    // Set Color to Bright Yellow
    SetConsoleTextAttribute(hConsole, FOREGROUND_RED | FOREGROUND_GREEN | FOREGROUND_INTENSITY);

    if (line > 0) {
        printf("[Warning] Line %d: %s\n", line, message);
    } else {
        printf("[Warning]: %s\n", message);
    }

    // Reset color
    SetConsoleTextAttribute(hConsole, saved_attributes);
}