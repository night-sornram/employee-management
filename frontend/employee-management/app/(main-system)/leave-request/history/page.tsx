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

export default function Page() {
    const { data: session } = useSession()
    const [date, setDate] = useState<Date>()
    const [data, setData] = useState<Leave[]>([]);
    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage  = 10
    const indexOfLastItem = currentPage * itemsPerPage;
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const [sort , setSort] = useState(true)
    const [selectedOption, setSelectedOption] = useState('all')
    let currentItems = data.slice(indexOfFirstItem, indexOfLastItem);

    const sortItem = (item : Leave[]) => {
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
        if(date === undefined){
        
            if(selectedOption === "month"){
                GetMyLeaves(session.user.employee_id, session.user.token).then((res) => {
                    setData(sortItem(res.filter((att : Leave) => {
                        return dayjs(att.date_start).local().format('MM/YYYY') === dayjs(date).local().format('MM/YYYY')
                    })));
                });
            }
            else if (selectedOption === "year"){
                GetMyLeaves(session.user.employee_id, session.user.token).then((res) => {
                    setData(sortItem(res.filter(
                        (att : Leave) => {
                            return dayjs(att.date_start).local().format('YYYY') === dayjs(date).local().format('YYYY')
                        }
                    )));
                });
            }
            else {
                GetMyLeaves(session.user.employee_id, session.user.token).then((res) => {
                    setData(sortItem(res));
                });
            }
        }
        else{
            GetMyLeaves(session.user.employee_id, session.user.token).then((res) => {
                setData(res.filter((att : Leave) => {
                    return dayjs(att.date_start).local().format('DD/MM/YYYY') === dayjs(date).local().format('DD/MM/YYYY')
                    || dayjs(att.date_end).local().format('DD/MM/YYYY') === dayjs(date).local().format('DD/MM/YYYY') 
                    || (new Date(att.date_end).getTime() > new Date(date).getTime() && new Date(att.date_start).getTime() < new Date(date).getTime())
                    }));
            });
        }
    }, [selectedOption , date , sort]);
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
                                <TableCell>
                                    {leave.category}
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
                                    <Link href={`/leave-request/history/${leave.id}`} className="hover:underline text-sky-600">
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
    )
}