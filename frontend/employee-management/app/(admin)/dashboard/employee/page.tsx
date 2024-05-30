"use client"

import * as React from "react"
import { CalendarIcon } from "@radix-ui/react-icons"
import { format } from "date-fns"
 
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Calendar } from "@/components/ui/calendar"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
  } from "@/components/ui/select"
import { TextAlignBottomIcon , TextAlignTopIcon} from '@radix-ui/react-icons'
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Attendance, UserJson } from "@/interface";
import getAllAttendances from "@/lib/GetAllAttendances";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import {
    Pagination,
    PaginationContent,
    PaginationEllipsis,
    PaginationItem,
    PaginationLink,
    PaginationNext,
    PaginationPrevious,
  } from "@/components/ui/pagination"
import { Input } from "@/components/ui/input";
import GetEmployee from "@/lib/GetEmployee";
import Link from "next/link"
  
dayjs.extend(utc);

export default function AllAttendanceHistoryPage () {

    const { data: session } = useSession()
    const [data, setData] = useState<UserJson[]>([]);
    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage  = 10
    const indexOfLastItem = currentPage * itemsPerPage;
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const [name, setName] = useState<string>("")
    let currentItems = data.slice(indexOfFirstItem, indexOfLastItem);
    const [staticData , setStaticData] = useState<UserJson[]>([])
    useEffect(() => {
        if (!session) {
            return () => {
                window.location.href = "/";
            };
        }
        GetEmployee(session.user.token).then((data) => {
            setStaticData(data);
            setData(data.filter(
                (user : UserJson) => user.role === "user"
            ));
        }
        ).catch((error) => {
            console.log(error);
        });
        
    }, [  ]);
    
    return (
        <main className='py-[3%] px-[5%]  md:w-[80%] 2xl:w-[60%] flex flex-col gap-10'>
            <div>
                <h1 className="font-bold text-2xl">
                    Employee
                </h1>
            </div>
            <div className="flex flex-row gap-10 md:overflow-y-hidden overflow-y-scroll">
                <Card className="w-[320px] ">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Employee
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {staticData.filter(
                            (user) => user.role === "user"
                        ).length}

                    </CardContent>
                </Card>
                
            </div>
            <div className=" flex flex-row space-x-3">
                <Input value={name} onChange={(e)=>{setName(e.currentTarget.value)}} type="text" placeholder="name"/>
            </div>
            <div className="">
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>
                                ID
                            </TableHead>
                            <TableHead>
                                Name
                            </TableHead>
                            <TableHead>
                                Department
                            </TableHead>
                            <TableHead>
                                Email
                            </TableHead>
                            <TableHead>
                                Phone
                            </TableHead>
                            <TableHead>
                                Detail
                            </TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            currentItems.filter(
                                (user) => {
                                    return( user.first_name_en + " " + user.last_name_en).toLowerCase().includes(name.toLowerCase()) ||
                                    (user.employee_id).toLowerCase().includes(name.toLowerCase()) ||
                                    (user.department).toLowerCase().includes(name.toLowerCase()) ||
                                    (user.email).toLowerCase().includes(name.toLowerCase()) ||
                                    (user.phone).toLowerCase().includes(name.toLowerCase())
                                }
                            ).map((user) => 
                            <TableRow >
                                <TableCell>
                                    {user.employee_id}
                                </TableCell>
                                <TableCell>
                                    {user.first_name_en + " " + user.last_name_en}
                                </TableCell>
                                <TableCell>
                                    {user.department}
                                </TableCell>
                                <TableCell>
                                    {user.email}
                                </TableCell>
                                <TableCell>
                                    {user.phone}
                                </TableCell>
                                <TableCell>                                    
                                    <Link href={`/dashboard/employee/${user.employee_id}`} className="hover:underline text-sky-600">
                                        Details
                                    </Link>                                   
                                </TableCell>
                            </TableRow>)
                        }
                    </TableBody>
                </Table>
            </div>
            <Pagination>
                <PaginationContent>
                    <PaginationItem>
                        <PaginationPrevious className=" cursor-pointer" onClick={()=>
                                {
                                    if(currentPage > 1){
                                        setCurrentPage(currentPage - 1)
                                    }}}
                            />
                    </PaginationItem>
                    <Input type="number" className=" w-10" value={currentPage} onChange={(e)=>
                        {
                            e.currentTarget.value === "" ? setCurrentPage(1) :
                            parseInt(e.currentTarget.value) > Math.ceil(data.length / itemsPerPage) ?
                            setCurrentPage(Math.ceil(data.length / itemsPerPage))
                            :
                            parseInt(e.currentTarget.value) < 1 ?
                            setCurrentPage(1)
                            :
                            setCurrentPage(parseInt(e.currentTarget.value))}
                        }
                    />
                    <input type="text" className=" w-10 text-center outline-none ring-0" value={"/  " + (Math.ceil(data.length / itemsPerPage) === 0 ? 1 : Math.ceil(data.length / itemsPerPage)) } readOnly/>

                    <PaginationItem>
                        <PaginationNext className=" cursor-pointer" onClick={()=>
                            {
                                if(currentPage < Math.ceil(data.length / itemsPerPage)){
                                    setCurrentPage(currentPage + 1)
                                }}
                            }
                            
                            />
                    </PaginationItem>
                </PaginationContent>
            </Pagination>
        </main>
    );
}