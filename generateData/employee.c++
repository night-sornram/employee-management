#include <iostream>
#include <string>
#include <iomanip>

using namespace std;

int main() {
    // Employee details
    string empId, thaiTitle, thaiFirstName, thaiLastName;
    string engTitle, engFirstName, engLastName;
    string dob, gender, department, position;
    string phone, email, password;

    for(int i = 0; i < 100; i++){
        string formattedEmpId = "EMP" + string(4 - (to_string(i)).length(), '0') + to_string(i);
        thaiTitle = "นาย";
        thaiFirstName = "thaiFirstName" + to_string(i);
        thaiLastName = "thaiLastName" + to_string(i);
        engTitle = "Mr.";
        engFirstName = "engFirstName"  + to_string(i);
        engLastName = "engLastName" + to_string(i);
        dob = "1980-01-01";
        gender = "Male";
        department = "IT";
        position = "user";
        phone = "0812345678";
        email = "uesr" + to_string(i) + "@example.com";
        password = "jZae727K08KaOmKSgOaGzww/XVqGr/PKEgIMkjrcbJI=";
        
        cout << "('" << formattedEmpId << "', '" << thaiTitle << "', '" << thaiFirstName << "', '" << thaiLastName << "', '"
         << engTitle << "', '" << engFirstName << "', '" << engLastName << "', '" << dob << "', '"
         << gender << "', '" << department << "', '" << position << "', '" << phone << "', '"
         << email << "', '" << password << "')," << endl;

    }
    return 0;
}
