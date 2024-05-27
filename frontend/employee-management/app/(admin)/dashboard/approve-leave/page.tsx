"use client"


import { Card, CardContent, CardHeader, CardDescription, CardTitle } from "@/components/ui/card";
import { Table, TableCaption, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Leave, UserJson } from "@/interface";
import GetLeaveAdmin from "@/lib/GetLeaveAdmin";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import { LapTimerIcon , CheckIcon ,Cross1Icon  } from "@radix-ui/react-icons";
import Link from "next/link";
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
    const [pending, setPending] = useState<Leave[]>([]);
    const [currentPagePending, setCurrentPagePending] = useState(1);
    const itemsPerPage  = 10
    const indexOfLastItemPending = currentPagePending * itemsPerPage;
    const indexOfFirstItemPending = indexOfLastItemPending - itemsPerPage;
    const currentItemsPending = pending.slice(indexOfFirstItemPending, indexOfLastItemPending);
    const [success, setSuccess] = useState<Leave[]>([]);
    const [currentPageSuccess, setCurrentPageSuccess] = useState(1);
    const indexOfLastItemSuccess = currentPageSuccess * itemsPerPage;
    const indexOfFirstItemSuccess = indexOfLastItemSuccess - itemsPerPage;
    const currentItemsSuccess = success.slice(indexOfFirstItemSuccess, indexOfLastItemSuccess);
    useEffect(() => {
        if (!session) {
            return () => {
                window.location.href = "/";
            };
        }
        GetLeaveAdmin(session.user.token).then((res) => {
            setPending(res.filter((leave : Leave) => leave.status == 'Pending').sort(function(a : any ,b : any){
                return Number(new Date(a.date_start)) - Number(new Date(b.date_start));
            }));
            setSuccess(res.filter((leave : Leave) => leave.status !== 'Pending').sort(function(a : any ,b : any){
                return Number(new Date(a.date_start)) - Number(new Date(b.date_start));
            }));
        });
    }, []);


    return(
        <main className='py-[5%] px-[5%] h-auto md:w-[70%] 2xl:w-[60%] flex flex-col gap-10'>
            <div>
                <h1 className="font-bold text-2xl">
                    Leave Management
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
                        {success.filter((leave) => leave.status == 'Approved').length}
                    </CardContent>
                </Card>
                <Card className="w-[320px]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Denied Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {success.filter((leave) => leave.status == 'Denied').length}
                    </CardContent>
                </Card>
                <Card className="w-[320px]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Pending Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {pending.length}
                    </CardContent>
                </Card>
            </div>
            <div className="">
                <h1 className="font-bold text-2xl">
                    Pending Request
                </h1>
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>
                                Employee ID
                            </TableHead>
                            <TableHead>
                                Employee
                            </TableHead>
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
                            <TableHead>
                                Details
                            </TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            pending.filter((leave) => leave.status == 'Pending').map((leave) =>
                            <TableRow  key={leave.id}>
                                <TableCell>
                                    {leave.employee_id}
                                </TableCell>
                                <TableCell>
                                    {leave.employee_name} {leave.employee_lastname}
                                </TableCell>
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
                                <TableCell>                                    
                                    <Link href={`/dashboard/approve-leave/${leave.id}`} className="hover:underline text-sky-600">
                                        Details
                                    </Link>                                   
                                </TableCell>
                            </TableRow>)
                        }
                    </TableBody>
                </Table>
                <Pagination  className=" pt-4 pb-10">
                    <PaginationContent>
                        <PaginationItem>
                            <PaginationPrevious className=" cursor-pointer" onClick={()=>
                                    {
                                        if(currentPagePending > 1){
                                            setCurrentPagePending(currentPagePending - 1)
                                        }}}
                                />
                        </PaginationItem>
                        <Input type="number" className=" w-10" value={currentPagePending} onChange={(e)=>
                            {
                                e.currentTarget.value === "" ? setCurrentPagePending(1) :
                                parseInt(e.currentTarget.value) > Math.ceil(pending.length / itemsPerPage) ?
                                setCurrentPagePending(Math.ceil(pending.length / itemsPerPage))
                                :
                                parseInt(e.currentTarget.value) < 1 ?
                                setCurrentPagePending(1)
                                :
                                setCurrentPagePending(parseInt(e.currentTarget.value))}
                            }
                        />
                        <input type="text" className=" w-10 text-center outline-none ring-0" value={"/  " + (Math.ceil(pending.length / itemsPerPage) === 0 ? 1 : Math.ceil(pending.length / itemsPerPage)) } readOnly/>

                        <PaginationItem>
                            <PaginationNext className=" cursor-pointer" onClick={()=>
                                {
                                    if(currentPagePending < Math.ceil(pending.length / itemsPerPage)){
                                        setCurrentPagePending(currentPagePending + 1)
                                    }}
                                }
                                
                                />
                        </PaginationItem>
                    </PaginationContent>
                </Pagination>

            </div>
            <div className="">
                <h1 className="font-bold text-2xl">
                    Success Request
                </h1>
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>
                                Employee ID
                            </TableHead>
                            <TableHead>
                                Employee
                            </TableHead>
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
                            <TableHead>
                                Details
                            </TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            success.filter((leave) => leave.status !== 'Pending').sort(
                                function(a,b){
                                    return Number(new Date(a.date_start)) - Number(new Date(b.date_start));
                                }
                            ).map((leave) => 
                            <TableRow key={leave.id}>
                                <TableCell>
                                    {leave.employee_id}
                                </TableCell>
                                <TableCell>
                                    {leave.employee_name} {leave.employee_lastname}
                                </TableCell>
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
                                <TableCell>                                    
                                    <Link href={`/dashboard/approve-leave/${leave.id}`} className="hover:underline text-sky-600">
                                        Details
                                    </Link>                                   
                                </TableCell>
                            </TableRow>)
                        }
                    </TableBody>
                </Table>
                <Pagination className=" pt-4">
                    <PaginationContent>
                        <PaginationItem>
                            <PaginationPrevious className=" cursor-pointer" onClick={()=>
                                    {
                                        if(currentPageSuccess > 1){
                                            setCurrentPageSuccess(currentPageSuccess - 1)
                                        }}}
                                />
                        </PaginationItem>
                        <Input type="number" className=" w-10" value={currentPageSuccess} onChange={(e)=>
                            {
                                e.currentTarget.value === "" ? setCurrentPageSuccess(1) :
                                parseInt(e.currentTarget.value) > Math.ceil(success.length / itemsPerPage) ?
                                setCurrentPageSuccess(Math.ceil(success.length / itemsPerPage))
                                :
                                parseInt(e.currentTarget.value) < 1 ?
                                setCurrentPageSuccess(1)
                                :
                                setCurrentPageSuccess(parseInt(e.currentTarget.value))}
                            }
                        />
                        <input type="text" className=" w-10 text-center outline-none ring-0" value={"/  " + (Math.ceil(success.length / itemsPerPage) === 0 ? 1 : Math.ceil(success.length / itemsPerPage)) } readOnly/>

                        <PaginationItem>
                            <PaginationNext className=" cursor-pointer" onClick={()=>
                                {
                                    if(currentPageSuccess < Math.ceil(success.length / itemsPerPage)){
                                        setCurrentPageSuccess(currentPageSuccess + 1)
                                    }}
                                }
                                
                                />
                        </PaginationItem>
                    </PaginationContent>
                </Pagination>
            </div>
        </main>
    )
}