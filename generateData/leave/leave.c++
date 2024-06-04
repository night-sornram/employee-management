#include <iostream>
#include <string>
#include <iomanip>

using namespace std;

int main() {
    for (int i = 0; i < 100; i++) {
        string formattedEmpId = "EMP" + string(4 - (to_string(i)).length(), '0') + to_string(i);
        string employeeID = formattedEmpId;
        string employeeName = "engFirstName" + to_string(i);
        string employeeLastName = "engLastName" + to_string(i);
        string dateStart = "2024-05-16T08:00:00Z";
        string dateEnd = "2024-05-17T18:00:00Z";
        string reason = "Reason" + to_string(i);
        string category = "Category" + to_string(i % 3); // Example categories
        string managerOpinion = "Opinion" + to_string(i);
        string status = "Pending";
        string manager = "Manager" + to_string(i % 5); // Example managers

        for (size_t j = 0; j < 2; j++) {
            cout << "('" << employeeID << "', '" << dateStart << "', '" << dateEnd << "', '"
                 << reason << "', '" << category << "', '" << managerOpinion << "', '" << status << "', '"
                 << manager << "', '" << employeeName << "', '" << employeeLastName << "')," << endl;
            dateStart = "2024-05-17T08:00:00Z";
            dateEnd = "2024-05-18T17:00:00Z";
            status = "approve";
        }
        dateStart = "0001-01-01T00:00:00Z";
        dateEnd = "0001-01-01T00:00:00Z";
        status = "denied";

        cout << "('" << employeeID << "', '" << dateStart << "', '" << dateEnd << "', '"
                << reason << "', '" << category << "', '" << managerOpinion << "', '" << status << "', '"
                << manager << "', '" << employeeName << "', '" << employeeLastName << "')," << endl;
    }
    return 0;
}
