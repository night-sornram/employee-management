"use client";

import { Button } from "@/components/ui/button";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
  } from "@/components/ui/select"
import { Label } from "@/components/ui/label";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import GetOneEmployee from "@/lib/GetOneEmployee";
import { UserJson } from "@/interface";


export default function Page({params} : {params: {id: string}}) {
    const { data: session } = useSession()
    const [loading, setLoading] = useState<boolean>(false)
    const [data, setData] = useState<UserJson>()

    useEffect(() => {
            if (!session) {
                return () => {
                    window.location.href = "/";
                };
            }
            GetOneEmployee(session.user.token, params.id).then((data) => {
                setData(data);
            }).catch((error) => {
                console.log(error);
            });

        }, []
    )

    return (
        <div className="flex flex-col md:space-y-7 space-y-5 px-[5%]  py-[5%] w-screen md:w-[80%] 2xl:w-[60%] gap-[5%]">
            <h1 className="text-2xl font-bold">Employee ID : {params.id}</h1>
            <div className="flex flex-row space-x-3 justify-between w-full">
                <div className=" w-2/12 flex flex-col space-y-3">
                    <Label htmlFor="title_eng">Title(ENG)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.title_en}
                    </div>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="name_eng">Name(ENG)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.first_name_en}
                    </div>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="surname_eng">Surname(ENG)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.last_name_en}
                    </div>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full">
                <div className=" w-2/12 flex flex-col space-y-3">
                    <Label htmlFor="title_th">Title(TH)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.title_th}
                    </div>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="name_th">Name(TH)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.first_name_th}
                    </div>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="surname_th">Surname(TH)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.last_name_th}
                    </div>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full">
            
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="date_of_birth">Date of birth</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.date_of_birth}
                    </div>
                </div>
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="surname_th">Gender</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.gender}
                    </div>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full"> 
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="department">Department</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.department}
                    </div>
                </div>
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="phone">Phone</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.phone}
                    </div>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full"> 
                <div className=" w-full flex flex-col space-y-3 ">
                    <Label htmlFor="email">Email</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {data?.email}
                    </div>
                </div>
            </div>
            
        </div>
    );
}