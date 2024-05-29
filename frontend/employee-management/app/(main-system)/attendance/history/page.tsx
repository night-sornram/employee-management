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
import Link from "next/link";
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

export default  function Page() {

    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage  = 10
    const indexOfLastItem = currentPage * itemsPerPage;
    const { data: session } = useSession()
    const [selectedOption, setSelectedOption] = useState('all')
    const [data, setData] = useState<Attendance[]>([]);
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const [date, setDate] = useState<Date>()
    const [sort , setSort] = useState(true)
    let currentItems = data.slice(indexOfFirstItem, indexOfLastItem);

    const sortItem = (item : Attendance[]) => {
        if(sort){
            return item.sort(function(a,b){
                return new Date(a.date).getTime() - new Date(b.date).getTime();
            });
        }
        else{
            return item.sort(function(a,b){
                
                return new Date(b.date).getTime() - new Date(a.date).getTime();
            });
        }
    }

    useEffect(() => {
        if (!session) {
            return () => {
                window.location.href = "/";
            };
        }
        if(date === undefined){
        
            if(selectedOption === "month"){
                GetMyAttendances(session.user.employee_id, session.user.token).then((res) => {
                    setData(sortItem(res.filter((att : Attendance) => {
                        return dayjs(att.date).local().format('MM/YYYY') === dayjs(date).local().format('MM/YYYY')
                    })));
                });
            }
            else if (selectedOption === "year"){
                GetMyAttendances(session.user.employee_id, session.user.token).then((res) => {
                    setData(sortItem(res.filter(
                        (att : Attendance) => {
                            return dayjs(att.date).local().format('YYYY') === dayjs(date).local().format('YYYY')
                        }
                    )));
                });
            }
            else {
                GetMyAttendances(session.user.employee_id, session.user.token).then((res) => {
                    setData(sortItem(res));
                });
            }
        }
        else{
            GetMyAttendances(session.user.employee_id, session.user.token).then((res) => {
                setData(res.filter((att : Attendance) => {
                    return dayjs(att.date).local().format('DD/MM/YYYY') === dayjs(date).local().format('DD/MM/YYYY')
                    }));
            });
        }
    }, [selectedOption , date , sort]);

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
            <div className=" w-full h-10 flex flex-row space-x-3">
                <Popover >
                    <PopoverTrigger asChild>
                        <Button
                        variant={"outline"}
                        className={cn(
                            "w-7/12 justify-start text-left font-normal",
                            !date && "text-muted-foreground"
                        )}
                        >
                        <CalendarIcon className="mr-2 h-4 w-4" />
                        {date ? format(date, "PPP") : <span>Pick a date</span>}
                        </Button>
                    </PopoverTrigger>
                    <PopoverContent className="w-auto p-0" align="start">
                        <Calendar
                        mode="single"
                        selected={date}
                        onSelect={setDate}
                        initialFocus
                        />
                    </PopoverContent>
                </Popover>
                <Select  value={selectedOption}
                onValueChange={(value) => {
                    setSelectedOption(value)
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
                    setSort(!sort)
                }}
                className=" w-2/12 flex justify-center items-center">
                    {
                        sort ?
                        <TextAlignBottomIcon className="h-5 w-5"/>
                        :
                        <TextAlignTopIcon className="h-5 w-5"/>
                    }
                </Button>
                <Button className=" w-1/12 flex justify-center items-center"
                onClick={
                    ()=>{
                        setDate(undefined)
                        setSelectedOption("all")
                        setSort(true)
                    }
                }>
                    Reset
                </Button>
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
                            <TableHead>
                                Leave
                            </TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            currentItems.map((att) => 
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
                                        att.leave_id !== -1?  "LEAVE"
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
                                <TableCell>
                                    {
                                        att.leave_id !== -1? 
                                        <Link href={`/leave-request/history/${att.leave_id}`} className="hover:underline text-sky-600">
                                            Details
                                        </Link>
                                        : "-"
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