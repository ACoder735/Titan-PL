#ifndef TITAN_TIME_H
#define TITAN_TIME_H

#include <time.h>
#include <windows.h>

TitanVar f_time_now(void) { 
    return titan_num((double)time(NULL)); 
}

TitanVar f_time_year(void) { 
    time_t t = time(NULL); 
    return titan_num(localtime(&t)->tm_year + 1900); 
}

TitanVar f_time_month(void) { 
    time_t t = time(NULL); 
    return titan_num(localtime(&t)->tm_mon + 1); 
}

TitanVar f_time_day(void) { 
    time_t t = time(NULL); 
    return titan_num(localtime(&t)->tm_mday); 
}

TitanVar f_time_hour(void) { 
    time_t t = time(NULL); 
    return titan_num(localtime(&t)->tm_hour); 
}

TitanVar f_time_minute(void) { 
    time_t t = time(NULL); 
    return titan_num(localtime(&t)->tm_min); 
}

TitanVar f_time_second(void) { 
    time_t t = time(NULL); 
    return titan_num(localtime(&t)->tm_sec); 
}

TitanVar f_time_weekday(void) { 
    time_t t = time(NULL); 
    return titan_num(localtime(&t)->tm_wday + 1);  // Add +1 for 1-7 range
}

TitanVar f_time_yearday(void) { 
    time_t t = time(NULL); 
    return titan_num(localtime(&t)->tm_yday); 
}

void f_time_sleep(TitanVar ms) { 
    Sleep((DWORD)ms.num); 
}

TitanVar f_time_clock(void) { 
    return titan_num((double)clock() / CLOCKS_PER_SEC); 
}

TitanVar f_time_timestamp(void) {
    SYSTEMTIME st;
    GetLocalTime(&st);
    char buf[32];
    sprintf(buf, "%04d-%02d-%02d %02d:%02d:%02d", 
            st.wYear, st.wMonth, st.wDay, 
            st.wHour, st.wMinute, st.wSecond);
    return titan_str(_strdup(buf));
}

TitanVar f_time_date(void) {
    time_t t = time(NULL);
    char* date = ctime(&t);
    date[24] = '\0';
    return titan_str(_strdup(date));
}

#endif