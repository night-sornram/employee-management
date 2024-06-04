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
dayjs.extend(utc);

export default function Page() {
    const { data: session } = useSession()
    const [pending, setPending] = useState<Leave[]>([]);
    const [currentPagePending, setCurrentPagePending] = useState(1);
    const itemsPerPage  = 10
    const indexOfLastItemPending = currentPagePending * itemsPerPage;
    const indexOfFirstItemPending = indexOfLastItemPending - itemsPerPage;
    const [date1, setDate1] = useState<Date>()
    const [sort1 , setSort1] = useState(true)
    const [selectedOption1, setSelectedOption1] = useState('all')
    let currentItemsPending = pending.slice(indexOfFirstItemPending, indexOfLastItemPending);
    const [success, setSuccess] = useState<Leave[]>([]);
    const [currentPageSuccess, setCurrentPageSuccess] = useState(1);
    const indexOfLastItemSuccess = currentPageSuccess * itemsPerPage;
    const indexOfFirstItemSuccess = indexOfLastItemSuccess - itemsPerPage;
    const [date2, setDate2] = useState<Date>()
    const [sort2 , setSort2] = useState(true)
    const [selectedOption2, setSelectedOption2] = useState('all')
    const currentItemsSuccess = success.slice(indexOfFirstItemSuccess, indexOfLastItemSuccess);
    const [name1, setName1] = useState('')
    const [name2, setName2] = useState('')

    const approved: string[] = ["approved", "Approved", "approve", "Approve"];
    const denied: string[] = ["denied", "Denied", "deny", "Deny"];

    const sortItem = (item : Leave[] , sort :boolean) => {
        if(sort){
            return item.sort(function(a,b){
                return new Date(a.date_start).getTime() - new Date(b.date_start).getTime();
            });
        }
        else{
            return item.sort(function(a,b){
                
                return new Date(b.date_start).getTime() - new Date(a.date_start).getTime();
            });
        }
    }

    useEffect(() => {
        if (!session) {
            return () => {
                window.location.href = "/";
            };
        }
        if(date1 === undefined){
        
            if(selectedOption1 === "month"){
                GetLeaveAdmin(session.user.token).then((res) => {
                    setPending(sortItem(res.filter((att : Leave) => {
                        return dayjs(att.date_start).local().format('MM/YYYY') === dayjs(date1).local().format('MM/YYYY')
                    }), sort1));
                });
            }
            else if (selectedOption1 === "year"){
                GetLeaveAdmin( session.user.token).then((res) => {
                    setPending(sortItem(res.filter(
                        (att : Leave) => {
                            return dayjs(att.date_start).local().format('YYYY') === dayjs(date1).local().format('YYYY')
                        }
                    ), sort1));
                });
            }
            else {
                GetLeaveAdmin( session.user.token).then((res) => {
                    setPending(sortItem(res , sort1));
                });
            }
        }
        else{
            GetLeaveAdmin(session.user.token).then((res) => {
                setPending(sortItem(
                    res.filter((att : Leave) => {
                        return dayjs(att.date_start).local().format('DD/MM/YYYY') === dayjs(date1).local().format('DD/MM/YYYY')
                        || dayjs(att.date_end).local().format('DD/MM/YYYY') === dayjs(date1).local().format('DD/MM/YYYY') 
                        || (new Date(att.date_end).getTime() > new Date(date1).getTime() && new Date(att.date_start).getTime() < new Date(date1).getTime())
                        }), sort1
                ));
            });
        }
        if(date2 === undefined){
        
            if(selectedOption2 === "month"){
                GetLeaveAdmin(session.user.token).then((res) => {
                    setSuccess(sortItem(res.filter((att : Leave) => {
                        return dayjs(att.date_start).local().format('MM/YYYY') === dayjs(date2).local().format('MM/YYYY')
                    }), sort2));
                });
            }
            else if (selectedOption2 === "year"){
                GetLeaveAdmin( session.user.token).then((res) => {
                    setSuccess(sortItem(res.filter(
                        (att : Leave) => {
                            return dayjs(att.date_start).local().format('YYYY') === dayjs(date2).local().format('YYYY')
                        }
                    ), sort2));
                });
            }
            else {
                GetLeaveAdmin( session.user.token).then((res) => {
                    setSuccess(sortItem(res , sort2));
                });
            }
        }
        else{
            GetLeaveAdmin(session.user.token).then((res) => {
                setSuccess(sortItem(
                    res.filter((att : Leave) => {
                        return dayjs(att.date_start).local().format('DD/MM/YYYY') === dayjs(date2).local().format('DD/MM/YYYY')
                        || dayjs(att.date_end).local().format('DD/MM/YYYY') === dayjs(date2).local().format('DD/MM/YYYY') 
                        || (new Date(att.date_end).getTime() > new Date(date2).getTime() && new Date(att.date_start).getTime() < new Date(date2).getTime())
                        }), sort2));
            });
        }
        console.log(success)
    }, [selectedOption1 , date1 , sort1 ,  selectedOption2 , date2 , sort2 ]);


    return(
        <main className='py-[5%] px-[5%] h-auto md:w-[80%] 2xl:w-[70%] flex flex-col gap-10'>
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
                        {success.filter((leave) => leave.status == "approve").length}
                    </CardContent>
                </Card>
                <Card className="w-[320px]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Denied Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {success.filter((leave) => leave.status == "denied").length}
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
            <div className=" flex flex-col space-y-5">
                <h1 className="font-bold text-2xl">
                    Pending Request
                </h1>
                <div className=" flex flex-col space-y-3">
                <Input value={name1} onChange={(e)=>{setName1(e.currentTarget.value)}} type="text" placeholder="name"/>
                <div className=" w-full h-10 flex flex-row space-x-3">
                    <Popover >
                        <PopoverTrigger asChild>
                            <Button
                            variant={"outline"}
                            className={cn(
                                "w-7/12 justify-start text-left font-normal",
                                !date1 && "text-muted-foreground"
                            )}
                            >
                            <CalendarIcon className="mr-2 h-4 w-4" />
                            {date1 ? format(date1, "PPP") : <span>Pick a date</span>}
                            </Button>
                            </PopoverTrigger>
                            <PopoverContent className="w-auto p-0" align="start">
                                <Calendar
                                mode="single"
                                selected={date1}
                                onSelect={setDate1}
                                initialFocus
                                />
                            </PopoverContent>
                        </Popover>
                        <Select  value={selectedOption1}
                        onValueChange={(value) => {
                            setSelectedOption1(value)
                        }}>
                            <SelectTrigger className="w-2/12">
                                <SelectValue placeholder="Select a fruit" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectGroup>
                                <SelectItem value="month">Month</SelectItem>
                                <SelectItem value="year">Year</SelectItem>
                                <SelectItem value="all">All</SelectItem>
                                </SelectGroup>
                            </SelectContent>
                        </Select>
                        <Button 
                        onClick={()=>{
                            setSort1(!sort1)
                        }}
                        className=" w-2/12 flex justify-center items-center">
                            {
                                sort1 ?
                                <TextAlignBottomIcon className="h-5 w-5"/>
                                :
                                <TextAlignTopIcon className="h-5 w-5"/>
                            }
                        </Button>
                        <Button className=" w-1/12 flex justify-center items-center"
                        onClick={
                            ()=>{
                                setDate1(undefined)
                                setSelectedOption1("all")
                                setSort1(true)
                                setName1("")
                            }
                        }>
                            Reset
                        </Button>
                    </div>
                </div>
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
                                Category
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
                            currentItemsPending.filter(
                                (leave) => {
                                    return (leave.employee_name + " " + leave.employee_lastname).toLowerCase().includes(name1.toLowerCase())
                                }
                            ).filter(
                                (leave) => {
                                    return leave.status === 'Pending' || leave.status === "pending";
                                }
                            ).map((leave) =>
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
                                <TableCell>
                                    {leave.category}
                                </TableCell>
                                {
                                    leave.status == "approve" ? 
                                    <TableCell className=" flex flex-row">
                                        <CheckIcon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell> : 
                                    denied.includes(leave.status) ?
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
            <div className=" flex flex-col space-y-5">
                <h1 className="font-bold text-2xl">
                    Success Request
                </h1>
                <div className=" flex flex-col space-y-3 ">
                <Input value={name2} onChange={(e)=>{setName2(e.currentTarget.value)}} type="text" placeholder="name"/>
                <div className=" w-full h-10 flex flex-row space-x-3">
                    <Popover >
                        <PopoverTrigger asChild>
                            <Button
                            variant={"outline"}
                            className={cn(
                                "w-7/12 justify-start text-left font-normal",
                                !date2 && "text-muted-foreground"
                            )}
                            >
                            <CalendarIcon className="mr-2 h-4 w-4" />
                            {date2 ? format(date2, "PPP") : <span>Pick a date</span>}
                            </Button>
                            </PopoverTrigger>
                            <PopoverContent className="w-auto p-0" align="start">
                                <Calendar
                                mode="single"
                                selected={date2}
                                onSelect={setDate2}
                                initialFocus
                                />
                            </PopoverContent>
                        </Popover>
                        <Select  value={selectedOption2}
                        onValueChange={(value) => {
                            setSelectedOption2(value)
                        }}>
                            <SelectTrigger className="w-2/12">
                                <SelectValue placeholder="Select a fruit" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectGroup>
                                <SelectItem value="month">Month</SelectItem>
                                <SelectItem value="year">Year</SelectItem>
                                <SelectItem value="all">All</SelectItem>
                                </SelectGroup>
                            </SelectContent>
                        </Select>
                        <Button 
                        onClick={()=>{
                            setSort2(!sort2)
                        }}
                        className=" w-2/12 flex justify-center items-center">
                            {
                                sort2 ?
                                <TextAlignBottomIcon className="h-5 w-5"/>
                                :
                                <TextAlignTopIcon className="h-5 w-5"/>
                            }
                        </Button>
                        <Button className=" w-1/12 flex justify-center items-center"
                        onClick={
                            ()=>{
                                setDate2(undefined)
                                setSelectedOption2("all")
                                setSort2(true)
                                setName2("")
                            }
                        }>
                            Reset
                        </Button>
                    </div>
                </div>
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
                                Category
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
                            currentItemsSuccess.filter(
                                (leave) => {
                                    return (leave.employee_name + " " + leave.employee_lastname).toLowerCase().includes(name2.toLowerCase())
                                }
                            ).filter(
                                (leave) => {
                                    return leave.status !== 'Pending'
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
                                <TableCell>
                                    {leave.category}
                                </TableCell>
                                {
                                    leave.status == 'approve' ? 
                                    <TableCell className=" flex flex-row">
                                        <CheckIcon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell> : 
                                    denied.includes(leave.status) ?
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