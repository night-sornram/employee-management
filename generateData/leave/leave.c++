#include <iostream>
#include <string>
#include <iomanip>

using namespace std;

int main() {
    int day = 5;
    int month = 3;
    int year = 2024;
    string categoryList[5] = {"Annual Leave", "Casual Leave", "Sick Leave", "Maternity Leave", "Unpaid Leave"};
    for (int i = 0; i < 100; i++) {
        string formattedEmpId = "EMP" + string(4 - (to_string(i)).length(), '0') + to_string(i);
        string employeeID = formattedEmpId;
        string employeeName = "";
        string employeeLastName = "";
        string dateStart;
        string dateEnd;
        string reason = "Reason" + to_string(i);
        string category = categoryList[i%5]; // Example categories
        string managerOpinion = "Opinion" + to_string(i);
        string status = "pending";
        string manager = "Manager" + to_string(i % 5); // Example managers

        dateStart = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-" + string(2 - (to_string(day)).length(), '0') + to_string(day) + "T01:00:00Z";
        dateEnd = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-" + string(2 - (to_string(day)).length(), '0') + to_string(day) + "T09:00:00Z";

        cout << "('" << employeeID << "', '" << dateStart << "', '" << dateEnd << "', '"
                << reason << "', '" << category << "', '" << managerOpinion << "', '" << status << "', '"
                << manager << "', '" << employeeName << "', '" << employeeLastName << "')," << endl;
        day++;

        if (day > 30) {
            day = 1;
            month++;
            if (month > 12) {
                month = 1;
                year++;
            }
        }

        status = "approved";
        dateStart = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-" + string(2 - (to_string(day)).length(), '0') + to_string(day) + "T01:00:00Z";
        dateEnd = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-" + string(2 - (to_string(day)).length(), '0') + to_string(day) + "T09:00:00Z";
        
        cout << "('" << employeeID << "', '" << dateStart << "', '" << dateEnd << "', '"
                        << reason << "', '" << category << "', '" << managerOpinion << "', '" << status << "', '"
                        << manager << "', '" << employeeName << "', '" << employeeLastName << "')," << endl;
        day++;

        if (day > 30) {
            day = 1;
            month++;
            if (month > 12) {
                month = 1;
                year++;
            }
        }
        dateStart = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-" + string(2 - (to_string(day)).length(), '0') + to_string(day) + "T01:00:00Z";
        dateEnd = to_string(year) + "-" + string(2 - (to_string(month)).length(), '0') + to_string(month) + "-" + string(2 - (to_string(day)).length(), '0') + to_string(day) + "T09:00:00Z";
        status = "denied";

        cout << "('" << employeeID << "', '" << dateStart << "', '" << dateEnd << "', '"
                << reason << "', '" << category << "', '" << managerOpinion << "', '" << status << "', '"
                << manager << "', '" << employeeName << "', '" << employeeLastName << "')," << endl;
        day++;
        if (day > 30) {
            day = 1;
            month++;
            if (month > 12) {
                month = 1;
                year++;
            }
        }
    }
    return 0;
}
