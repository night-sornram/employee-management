'use client'
import UserCard from "@/components/User/UserCard";
import { UserJson } from "@/interface";
import dayjs, { Dayjs } from "dayjs";
import utc from "dayjs/plugin/utc";
import { useState } from "react";
dayjs.extend(utc);

export default function AttendancePage () {

    const userData:UserJson = {
        name: "Prasut",
        lastname: "Parinippan",
        title: "Mr",
        gender: "Male",
        role: "Employee"
    }
    const date = dayjs().utc().local().format('DD/MM/YYYY');
    const time = dayjs().utc().local().format('HH:mm:ss');

    const [checkedIn, setCheckedIn] = useState(false);
    const [checkedOut, setCheckedOut] = useState(false);

    return (
        <div className="w-[85%] pt-[5%] px-[5%] flex flex-col gap-[5%]">
            <UserCard userData={userData}/>
            <h1 className="text-2xl">Date: {date}</h1>
            <table className="shadow-md table-fixed">
                <thead>
                    <tr className="bg-p-gray-2">
                        <th className="py-4 font-semibold w-[50%]">
                                Check-in
                        </th>
                        <th className="py-4 font-semibold w-[50%]">
                                Check-out
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr className="bg-p-lightgray py-5">
                        <td className="text-center py-4 w-[50%] h-20">
                            {
                                checkedIn? time : 
                                <button className="custom-btn-green hover:custom-btn-green-hover"
                                onClick={(e) => setCheckedIn(true)}>
                                    Check-in
                                </button>
                            }
                            
                        </td>
                        <td className="text-center py-4 w-[50%] h-20">
                            {
                                (!checkedIn)? 
                                <button className="custom-btn-dark">
                                    Check-out
                                </button> : (checkedOut)?
                                time :
                                <button className="custom-btn-green hover:custom-btn-green-hover"
                                onClick={(e) => setCheckedOut(true)}>
                                    Check-out
                                </button>
                            }
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    );
}