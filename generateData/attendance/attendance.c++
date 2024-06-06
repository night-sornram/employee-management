#include <iostream>
#include <string>
#include <iomanip>

using namespace std;

int main() {
    int day = 1;
    int month = 3;
    int year = 2024;

    for(int i = 0; i < 100; i++){
        string formattedEmpId = "EMP" + string(4 - (to_string(i)).length(), '0') + to_string(i);
        int leaveID = -1;
        string date = "";
        string employeeName = "";
        string employeeLastName = "";

        for (size_t j = 0; j < 2; j++)
        {
            if (day > 29) {
                day = 1;
                month++;
                if (month > 12) {
                    month = 1;
                    year++;
                }
            }

            date = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-"  + string(2 - (to_string(day)).length(), '0') + to_string(day) ;
            string checkIn = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-" + string(2 - (to_string(day)).length(), '0') + to_string(day) + "T01:00:00Z";
            string checkOut = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-" + string(2 - (to_string(day)).length(), '0') + to_string(day) + "T09:00:00Z";
            cout << "('" << formattedEmpId << "', '" << checkIn << "', '" << checkOut << "', '" << date << "', '"
            << leaveID << "', '" << employeeName << "', '" << employeeLastName << "')," << endl;
            day++;
        }

        date = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-"  + string(2 - (to_string(day)).length(), '0') + to_string(day) ;
        string checkIn = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-" + string(2 - (to_string(day)).length(), '0') + to_string(day) + "T04:00:00Z";
        string checkOut = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-" + string(2 - (to_string(day)).length(), '0') + to_string(day) + "T12:00:00Z";
        cout << "('" << formattedEmpId << "', '" << checkIn << "', '" << checkOut << "', '" << date << "', '"
        << leaveID << "', '" << employeeName << "', '" << employeeLastName << "')," << endl;
        day++;

        leaveID = i + 1;
        checkIn = "0001-01-01T00:00:00Z";
        checkOut = "0001-01-01T00:00:00Z";
        cout << "('" << formattedEmpId << "', '" << checkIn << "', '" << checkOut << "', '" << date << "', '"
            << leaveID << "', '" << employeeName << "', '" << employeeLastName << "')," << endl;
        day -= 1;
    }
    return 0;
}
