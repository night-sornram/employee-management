#include <iostream>
#include <string>
#include <iomanip>

using namespace std;

int main() {
    string departmentList[6] = {"R&D", "IT", "Software Development", "Product Management", "QA", "HR"};
    string genderList[2] = {"Male", "Female"};
    for(int i = 0; i < 100; i++){
        string formattedEmpId = "EMP" + string(4 - (to_string(i)).length(), '0') + to_string(i);
        string thaiTitle = "นาย";
        string thaiFirstName = "thaiFirstName_" + to_string(i);
        string thaiLastName = "thaiLastName_" + to_string(i);
        string engTitle = "Mr.";
        string engFirstName = "engFirstName_"  + to_string(i);
        string engLastName = "engLastName_" + to_string(i);
        string dob = "1980-01-01";
        string gender = genderList[i%2];
        string department = departmentList[i%6];
        string position = "user";
        string phone = "0812345678";
        string email = "uesr" + to_string(i) + "@example.com";
        string password = "jZae727K08KaOmKSgOaGzww/XVqGr/PKEgIMkjrcbJI=";
        
        cout << "('" << formattedEmpId << "', '" << thaiTitle << "', '" << thaiFirstName << "', '" << thaiLastName << "', '"
         << engTitle << "', '" << engFirstName << "', '" << engLastName << "', '" << dob << "', '"
         << gender << "', '" << department << "', '" << position << "', '" << phone << "', '"
         << email << "', '" << password << "')," << endl;

    }
    return 0;
}
