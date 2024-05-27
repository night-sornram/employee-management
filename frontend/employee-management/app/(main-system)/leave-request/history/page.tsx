"use client"

import { Card, CardContent, CardHeader, CardDescription, CardTitle } from "@/components/ui/card";
import { Table, TableCaption, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Leave, UserJson } from "@/interface";
import GetMyLeaves from "@/lib/GetMyLeaves";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import { LapTimerIcon , CheckIcon ,Cross1Icon  } from "@radix-ui/react-icons";
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
dayjs.extend(utc);

export default function Page() {
    const { data: session } = useSession()
    const [data, setData] = useState<Leave[]>([]);
    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage  = 10
    const indexOfLastItem = currentPage * itemsPerPage;
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const currentItems = data.slice(indexOfFirstItem, indexOfLastItem);


    useEffect(() => {
        if (!session) {
            return () => {
                window.location.href = "/";
            };
        }
        GetMyLeaves(session.user.employee_id, session.user.token).then((res) => {
            setData(res.sort(function(a : any ,b : any){
                return Number(new Date(a.date_start)) - Number(new Date(b.date_start));
            }));
        });
    }, []);
    return(
        <main className='py-[3%] px-[5%] h-full  md:w-[80%] 2xl:w-[60%] flex flex-col gap-10'>
            <div>
                <h1 className="font-bold text-2xl">
                    History of Leave
                </h1>
            </div>
            <div className="flex flex-row gap-10 md:overflow-y-hidden overflow-y-scroll">
                <Card className="w-[320px]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Approved Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {data.filter((leave) => leave.status == 'Approved').length}
                    </CardContent>
                </Card>
                <Card className="w-[320px]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Denied Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {data.filter((leave) => leave.status == 'Denied').length}
                    </CardContent>
                </Card>
                <Card className="w-[320px]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Pending Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {data.filter((leave) => leave.status == 'Pending').length}
                    </CardContent>
                </Card>
            </div>
            <div className="">
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>
                                From
                            </TableHead>
                            <TableHead>
                                To
                            </TableHead>
                            <TableHead>
                                Duration (days)
                            </TableHead>
                            <TableHead>
                                Status
                            </TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            currentItems.map((leave) => 
                            <TableRow key={leave.id}>
                                <TableCell>
                                    {dayjs(leave.date_start).format('DD/MM/YYYY')}
                                </TableCell>
                                <TableCell>
                                    {dayjs(leave.date_end).format('DD/MM/YYYY')}
                                </TableCell>
                                <TableCell>
                                    {dayjs(leave.date_end).diff(dayjs(leave.date_start), 'day') + 1}
                                </TableCell>
                                {
                                    leave.status == "Approved" ? 
                                    <TableCell className=" flex flex-row">
                                        <CheckIcon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell> : 
                                    leave.status == "Denied" ?
                                    <TableCell className=" flex flex-row">
                                        <Cross1Icon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell> :
                                    <TableCell className=" flex flex-row">
                                        <LapTimerIcon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell>
                                }
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