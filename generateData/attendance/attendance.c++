#include <iostream>
#include <string>
#include <iomanip>

using namespace std;

int main() {
    for(int i = 0; i < 100; i++){
        string formattedEmpId = "EMP" + string(4 - (to_string(i)).length(), '0') + to_string(i);
        string checkIn = "2024-05-14T08:00:00Z";
        string checkOut = "2024-05-14T17:00:00Z";
        string date = "14/05/2024";
        int leaveID = -1;
        string employeeName = "engFirstName"  + to_string(i);
        string employeeLastName = "engLastName" + to_string(i);
        for (size_t j = 0; j < 2; j++)
        {
            cout << "('" << formattedEmpId << "', '" << checkIn << "', '" << checkOut << "', '" << date << "', '"
            << leaveID << "', '" << employeeName << "', '" << employeeLastName << "')," << endl;
            checkIn = "2024-05-15T011:00:00Z";
            checkOut = "2024-05-15T17:00:00Z";
        }
        leaveID = i + 1;
        checkIn = "0001-01-01T00:00:00Z";
        checkOut = "0001-01-01T00:00:00Z";
        cout << "('" << formattedEmpId << "', '" << checkIn << "', '" << checkOut << "', '" << date << "', '"
            << leaveID << "', '" << employeeName << "', '" << employeeLastName << "')," << endl;

    }
    return 0;
}
