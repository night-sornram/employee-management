"use client"

import { Card, CardContent, CardHeader, CardDescription, CardTitle } from "@/components/ui/card";
import { Table, TableCaption, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Attendance, UserJson } from "@/interface";
import GetMyAttendances from "@/lib/GetMyAttendances";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import { useEffect, useState } from "react";
import { useSession } from "next-auth/react";
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
dayjs.extend(utc);

export default  function Page() {

    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage  = 10
    const indexOfLastItem = currentPage * itemsPerPage;
    const { data: session } = useSession()

    const [data, setData] = useState<Attendance[]>([]);
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const currentItems = data.slice(indexOfFirstItem, indexOfLastItem);

    useEffect(() => {
        if (!session) {
            return () => {
                window.location.href = "/";
            };
        }
        GetMyAttendances(session.user.employee_id, session.user.token).then((res) => {
            setData(res.sort(function(a : any ,b : any){
                return Number(new Date(a.date)) - Number(new Date(b.date));
            }));
        });
    }, []);

    const countLeave = (data: Attendance) => {
        return data.leave_id != -1;
    }

    return(
        <main className='py-[3%] px-[5%] h-full  md:w-[80%] 2xl:w-[60%] flex flex-col gap-10'>
            <div>
                <h1 className="font-bold text-2xl">
                    History of Attendance
                </h1>
            </div>
            <div className="flex flex-row gap-10  md:overflow-y-hidden overflow-y-scroll">
                <Card className="w-[320px] ">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            All-time Attendance
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {data.length - data.filter(countLeave).length}
                    </CardContent>
                </Card>
                <Card className="w-[320px]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            All-time Absence
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {data.filter(countLeave).length}
                    </CardContent>
                </Card>
            </div>
            <div className="">
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>
                                Date
                            </TableHead>
                            <TableHead>
                                Check-in Time
                            </TableHead>
                            <TableHead>
                                Check-out Time
                            </TableHead>
                            <TableHead>
                                Duration (Hour)
                            </TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            currentItems.sort(function(a,b){
                                return Number(new Date(a.date)) - Number(new Date(b.date));
                            }).map((att) => 
                            <TableRow key={att.id}>
                                <TableCell>
                                    {dayjs(att.date).local().format('DD/MM/YYYY')}
                                </TableCell>
                                <TableCell>
                                    {att.leave_id !== -1 ? "LEAVE" : 
                                        (
                                            dayjs(att.check_in).local().format('HH:mm:ss')
                                        )
                                    }
                                    
                                </TableCell>
                                <TableCell>
                                    {   att.leave_id !== -1? "LEAVE" :
                                        (
                                            dayjs(att.check_out).local().toString() === "Mon, 01 Jan 0001 00:00:00 GMT" ?
                                            "-" :
                                            (
                                                dayjs(att.check_out).local().format('HH:mm:ss')
                                            )
                                        )
                                    }
                                    
                                </TableCell>
                                <TableCell>
                                    {
                                        att.leave_id !== -1? "LEAVE" 
                                        : 
                                        (
                                            dayjs(att.check_out).local().toString() === "Mon, 01 Jan 0001 00:00:00 GMT" ?
                                            "-" :
                                            (
                                                (Math.round(dayjs(att.check_out).diff(dayjs(att.check_in), 'hour', true) * 100) / 100).toFixed(2) +  " Hrs" 
                                            )
                                        )
                                    } 
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
    )
}