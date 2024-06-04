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
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat);
dayjs.extend(utc);

export default function AllAttendanceHistoryPage () {

    const { data: session } = useSession()
    const [data, setData] = useState<Attendance[]>([]);
    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage  = 10
    const indexOfLastItem = currentPage * itemsPerPage;
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const [selectedOption, setSelectedOption] = useState('all')
    const [date, setDate] = useState<Date>()
    const [sort , setSort] = useState(true)
    const [name, setName] = useState<string>("")
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
                getAllAttendances( session.user.token).then((res) => {
                    setData(sortItem(res.filter((att : Attendance) => {
                        return dayjs(att.date).local().format('MM/YYYY') === dayjs(date).local().format('MM/YYYY')
                    })));
                });
            }
            else if (selectedOption === "year"){
                getAllAttendances( session.user.token).then((res) => {
                    setData(sortItem(res.filter(
                        (att : Attendance) => {
                            return dayjs(att.date).local().format('YYYY') === dayjs(date).local().format('YYYY')
                        }
                    )));
                });
            }
            else {
                getAllAttendances( session.user.token).then((res) => {
                    setData(sortItem(res));
                });
            }
        }
        else{
            getAllAttendances( session.user.token).then((res) => {
                setData(res.filter((att : Attendance) => {
                    return dayjs(att.date).local().format('DD/MM/YYYY') === dayjs(date).local().format('DD/MM/YYYY')
                    }));
            });
        }
    }, [selectedOption , sort, date ]);
    
    return (
        <main className='py-[5%] px-[5%]  md:w-[80%] 2xl:w-[60%] flex flex-col gap-10'>
            <div>
                <h1 className="font-bold text-2xl">
                    Employee Attendance
                </h1>
            </div>
            <div className="flex flex-row gap-10 md:overflow-y-hidden overflow-y-scroll">
                <Card className="w-[320px] ">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            All-time Attendance
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {data.length - data.filter((att) => att.leave_id !== -1).length}
                    </CardContent>
                </Card>
                <Card className="w-[320px]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            All-time Absence
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {data.filter((att) => att.leave_id !== -1).length}
                    </CardContent>
                </Card>
            </div>
            <div className=" flex flex-col space-y-3">
                <Input value={name} onChange={(e)=>{setName(e.currentTarget.value)}} type="text" placeholder="name"/>
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
                            setName("")
                        }
                    }>
                        Reset
                    </Button>
                </div>
            </div>
            <div className="">
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>
                                Date
                            </TableHead>
                            <TableHead>
                                Employee
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
                            currentItems.filter(
                                (att) => {
                                    return( att.employee_name + " " + att.employee_lastname).toLowerCase().includes(name.toLowerCase())
                                }
                            ).map((att) => 
                            <TableRow key={att.id}>
                                <TableCell>
                                    {
                                        dayjs(att.date, ["DD/MM/YYYY", "DD/MM/YY", "DD-MM-YY", "DD-MM-YYYY", "YYYY-MM-DD"]).local().format('DD/MM/YYYY')   
                                    }
                                </TableCell>
                                <TableCell>
                                    {att.employee_name} {att.employee_lastname}
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
    );
}