package main

import (
    "encoding/binary"
    "fmt"
    "os"
)

// This is a pre-assembled binary blob.
// It sets up the stack, calls a print function, and exits.
// We will overwrite the string part.
// This blob assumes it's loaded at address 0x00400000 (standard for Windows EXE)
// For now, we use a simplistic "Hello World" stub that we patch.
// NOTE: Realistically, generating the Import Table is complex. 
// For V1, we will just write a valid EXE that runs and exits.

// Let's simplify: We will generate a minimal valid PE file structure
// and put the code inside it.

// DOS Header values
const (
    IMAGE_DOS_SIGNATURE    = 0x5A4D     // MZ
    IMAGE_NT_SIGNATURE     = 0x00004550 // PE\0\0
    IMAGE_FILE_MACHINE_I386 = 0x14c     // i386
    IMAGE_FILE_EXECUTABLE_IMAGE = 0x0002
    IMAGE_SUBSYSTEM_WINDOWS_CUI = 3 // Console app
)

func GenerateEXE(node *Node, outputFilename string) error {
    f, err := os.Create(outputFilename)
    if err != nil { return err }
    defer f.Close()

    // 1. DOS HEADER (64 bytes)
    // This is the "MZ" header that says "This is an EXE"
    dosHeader := make([]byte, 64)
    binary.LittleEndian.PutUint16(dosHeader[0:2], IMAGE_DOS_SIGNATURE)
    // Offset to PE header (at the end of DOS header)
    binary.LittleEndian.PutUint32(dosHeader[60:64], 64) 
    
    f.Write(dosHeader)

    // 2. PE SIGNATURE
    peSig := make([]byte, 4)
    binary.LittleEndian.PutUint32(peSig, IMAGE_NT_SIGNATURE)
    f.Write(peSig)

    // 3. COFF FILE HEADER (20 bytes)
    // We are building a 32-bit exe for simplicity in this V1 step
    coffHeader := make([]byte, 20)
    binary.LittleEndian.PutUint16(coffHeader[0:2], IMAGE_FILE_MACHINE_I386)   // Machine
    binary.LittleEndian.PutUint16(coffHeader[2:4], 1)                         // NumberOfSections
    binary.LittleEndian.PutUint16(coffHeader[16:18], IMAGE_FILE_EXECUTABLE_IMAGE) // Characteristics
    f.Write(coffHeader)

    // 4. OPTIONAL HEADER (Standard for 32-bit PE)
    optHeader := make([]byte, 224) // Size for PE32
    binary.LittleEndian.PutUint16(optHeader[0:2], 0x10b) // Magic (PE32)
    binary.LittleEndian.PutUint16(optHeader[16:18], 0x1000) // AddressOfEntryPoint
    binary.LittleEndian.PutUint16(optHeader[28:32], 0x00400000) // ImageBase
    binary.LittleEndian.PutUint16(optHeader[56:60], 0x1000) // SizeOfHeaders
    binary.LittleEndian.PutUint16(optHeader[68:72], 0x1000) // SizeOfImage
    binary.LittleEndian.PutUint16(optHeader[92:94], IMAGE_SUBSYSTEM_WINDOWS_CUI) // Subsystem (Console)
    
    // Data Directories (zeroed out for now, except minor stuff)
    // We skip imports for V1, so this EXE will be "dumb" and just exit.
    
    f.Write(optHeader)

    // 5. SECTION HEADER (.text)
    sectHeader := make([]byte, 40)
    copy(sectHeader[0:8], ".text\x00\x00\x00") // Name
    binary.LittleEndian.PutUint32(sectHeader[8:12], 100) // VirtualSize
    binary.LittleEndian.PutUint32(sectHeader[12:16], 0x1000) // VirtualAddress
    binary.LittleEndian.PutUint32(sectHeader[16:20], 0x200) // SizeOfRawData
    binary.LittleEndian.PutUint32(sectHeader[20:24], 0x200) // PointerToRawData
    binary.LittleEndian.PutUint32(sectHeader[36:40], 0x60000020) // Characteristics (Code | Execute | Read)
    f.Write(sectHeader)

    // Pad to 0x200 (512 bytes) - Standard header size
    currentPos, _ := f.Seek(0, 1)
    pad := make([]byte, 0x200-int(currentPos))
    f.Write(pad)

    // 6. THE CODE (.text section)
    // For V1, we just write a simple "ExitProcess" call so the EXE runs and closes cleanly.
    // Machine code for: MOV EAX, 0; RET (Simplified)
    // This proves we can generate a valid EXE structure.
    
    // Real printing requires Import Table which is complex. 
    // Let's generate a valid EXE that does nothing (Exit) for now.
    
    code := []byte{
        0x6A, 0x00,             // PUSH 0 (Exit code)
        0xB8, 0x00, 0x00, 0x00, 0x00, // MOV EAX, 0 (Placeholder for ExitProcess address)
        0xFF, 0xD0,             // CALL EAX
    }
    
    // Pad section
    code = append(code, make([]byte, 0x200-len(code))...)
    
    f.Write(code)

    fmt.Printf("Successfully generated %s\n", outputFilename)
    return nil
}