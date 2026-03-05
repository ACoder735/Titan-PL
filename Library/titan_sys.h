#ifndef TITAN_SYS_H
#define TITAN_SYS_H

extern int _titan_argc;
extern char** _titan_argv;

TitanVar f_sys_args(void) {
    TitanVar arr = titan_new_array();
    for (int i = 0; i < _titan_argc; i++) {
        titan_arr_push(arr, titan_str(_titan_argv[i]));
    }
    return arr;
}

void f_sys_exit(TitanVar code) { 
    exit((int)code.num); 
}

TitanVar f_sys_os(void) { 
    return titan_str("windows"); 
}

TitanVar f_sys_arch(void) {
#if defined(_M_X64) || defined(__x86_64__)
    return titan_str("x64");
#elif defined(_M_IX86) || defined(__i386__)
    return titan_str("x86");
#elif defined(_M_ARM64) || defined(__aarch64__)
    return titan_str("arm64");
#else
    return titan_str("unknown");
#endif
}

void f_sys_pause(void) { 
    system("pause"); 
}

TitanVar f_sys_env(TitanVar name) {
    if (name.type != 2) return TITAN_NULL;
    char* val = getenv(name.str);
    if (val) return titan_str(_strdup(val));
    return TITAN_NULL;
}

TitanVar f_sys_cwd(void) {
    char buf[1024];
    if (GetCurrentDirectoryA(sizeof(buf), buf)) {
        return titan_str(_strdup(buf));
    }
    return TITAN_NULL;
}

void f_sys_cd(TitanVar path) {
    if (path.type == 2) SetCurrentDirectoryA(path.str);
}

TitanVar f_sys_exec(TitanVar cmd) {
    if (cmd.type != 2) return titan_num(-1);
    return titan_num((double)system(cmd.str));
}

#endif