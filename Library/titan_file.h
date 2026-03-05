#ifndef TITAN_FILE_H
#define TITAN_FILE_H

#include <stdio.h>
#include <stdlib.h>

TitanVar f_file_read(TitanVar path) {
    if (path.type != 2) return TITAN_NULL;
    FILE* f = fopen(path.str, "r");
    if (!f) return TITAN_NULL;
    fseek(f, 0, SEEK_END);
    long len = ftell(f);
    fseek(f, 0, SEEK_SET);
    char* buf = malloc(len + 1);
    if (!buf) { fclose(f); return TITAN_NULL; }
    fread(buf, 1, len, f);
    buf[len] = '\0';
    fclose(f);
    return titan_str(buf);
}

void f_file_write(TitanVar path, TitanVar content) {
    if (path.type != 2 || content.type != 2) return;
    FILE* f = fopen(path.str, "w");
    if (f) { fputs(content.str, f); fclose(f); }
}

void f_file_append(TitanVar path, TitanVar content) {
    if (path.type != 2 || content.type != 2) return;
    FILE* f = fopen(path.str, "a");
    if (f) { fputs(content.str, f); fclose(f); }
}

TitanVar f_file_exists(TitanVar path) {
    if (path.type != 2) return titan_num(0);
    FILE* f = fopen(path.str, "r");
    if (f) { fclose(f); return titan_num(1); }
    return titan_num(0);
}

void f_file_delete(TitanVar path) {
    if (path.type == 2) remove(path.str);
}

TitanVar f_file_size(TitanVar path) {
    if (path.type != 2) return titan_num(-1);
    FILE* f = fopen(path.str, "rb");
    if (!f) return titan_num(-1);
    fseek(f, 0, SEEK_END);
    long size = ftell(f);
    fclose(f);
    return titan_num((double)size);
}

TitanVar f_file_copy(TitanVar src, TitanVar dest) {
    if (src.type != 2 || dest.type != 2) return titan_num(0);
    FILE* fsrc = fopen(src.str, "rb");
    if (!fsrc) return titan_num(0);
    FILE* fdest = fopen(dest.str, "wb");
    if (!fdest) { fclose(fsrc); return titan_num(0); }
    char buf[4096];
    size_t n;
    while ((n = fread(buf, 1, sizeof(buf), fsrc)) > 0) {
        fwrite(buf, 1, n, fdest);
    }
    fclose(fsrc);
    fclose(fdest);
    return titan_num(1);
}

TitanVar f_file_rename(TitanVar oldname, TitanVar newname) {
    if (oldname.type != 2 || newname.type != 2) return titan_num(0);
    return titan_num(rename(oldname.str, newname.str) == 0 ? 1 : 0);
}

#endif