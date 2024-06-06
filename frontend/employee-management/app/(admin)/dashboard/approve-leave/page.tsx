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
import { DataJson, Leave, UserJson } from "@/interface";
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
import getAllAttendances from "@/lib/GetAllAttendances"
import DownloadLeave from "@/lib/DownloadLeave"
dayjs.extend(utc);

export default function Page() {
    const { data: session } = useSession()
    const [pending, setPending] = useState<Leave[]>([]);
    const [currentPagePending, setCurrentPagePending] = useState(1);
    const [date1, setDate1] = useState<Date>()
    const [sort1 , setSort1] = useState(true)
    const [selectedOption1, setSelectedOption1] = useState('all')
    const [success, setSuccess] = useState<Leave[]>([]);
    const [currentPageSuccess, setCurrentPageSuccess] = useState(1);
    const [date2, setDate2] = useState<Date>()
    const [sort2 , setSort2] = useState(true)
    const [selectedOption2, setSelectedOption2] = useState('all')
    const [name1, setName1] = useState('')
    const [name2, setName2] = useState('')
    const [json1, setJson1] = useState<DataJson>()
    const [json2, setJson2] = useState<DataJson>()

    const [downloadUrl, setDownloadUrl] = useState<string | null>();

    const handleDownload = async () => {
        if (!session) return;
        
        try {
            const blob = await DownloadLeave(session.user.token);
            const url = URL.createObjectURL(blob);
            setDownloadUrl(url);
        } catch (error) {
            console.error("Error downloading attendance data:", error);
        }
    };

    const sortItem = (item : Leave[] , sort :boolean) => {
        if ( item === null) return [];
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
        let query1 : string = `?page=${currentPagePending}&status=pending`
        
        if(date1){
            let year = date1.toLocaleString("default", { year: "numeric" });
            let month = date1.toLocaleString("default", { month: "2-digit" });
            let day = date1.toLocaleString("default", { day: "2-digit" });

            let formattedDate = year + "-" + month + "-" + day;
            query1 += `&date=${formattedDate}`
        }
        if (name1){
            query1 += `&name=${name1}`
        }
        if (selectedOption1 === "all"){
            query1 += `&option=All`
        }
        else if (selectedOption1 === "year"){
            query1 += `&option=Year`
        }
        else{
            query1 += `&option=Month`
        }
        console.log(query1)
        GetLeaveAdmin( session.user.token , query1).then((res) => {
            setPending(
                sortItem(res.data , sort1)
            );
            setJson1(res)
        })

        let query2 : string = `?page=${currentPageSuccess}&status=approved`
        
        if(date2){
            let year = date2.toLocaleString("default", { year: "numeric" });
            let month = date2.toLocaleString("default", { month: "2-digit" });
            let day = date2.toLocaleString("default", { day: "2-digit" });

            let formattedDate = year + "-" + month + "-" + day;
            query2 += `&date=${formattedDate}`
        }

        if (name2){
            query2 += `&name=${name2}`
        }
        if (selectedOption2 === "all"){
            query2 += `&option=All`
        }
        else if (selectedOption2 === "year"){
            query2 += `&option=Year`
        }
        else{
            query2 += `&option=Month`
        }
        
        GetLeaveAdmin( session.user.token , query2).then((res) => {
            setSuccess(
                sortItem(res.data , sort2)
            );
            setJson2(res)
        })
    }, [selectedOption1 , date1 , sort1, currentPageSuccess,selectedOption2 , date2 , sort2, currentPagePending , name1 , name2]);

    useEffect(() => {
        handleDownload();
    }, []);

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
                            Success Request
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {json2?.total}
                    </CardContent>
                </Card>
                <Card className="w-[320px]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Pending Request
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {json1?.total}
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
                            pending.map((leave) =>
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
                                    leave.status == "approved" ? 
                                    <TableCell className=" flex flex-row">
                                        <CheckIcon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell> : 
                                    leave.status == "denied" ?
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
                <Pagination className=" pt-4 pb-10">
                <PaginationContent>
                    <PaginationItem>
                        <PaginationPrevious className=" cursor-pointer" onClick={()=>
                                {
                                    if(json1 && json1.page > 1){
                                        setCurrentPagePending(currentPagePending - 1)
                                    }}}
                            />
                    </PaginationItem>
                    <Input type="number" className=" w-10" value={currentPagePending} onChange={(e)=>
                        {
                            if (json1){
                                if (e.currentTarget.value === ""){
                                    setCurrentPagePending(1)
                                }
                                else if(parseInt(e.currentTarget.value) > json1.last_page){
                                    setCurrentPagePending(json1.last_page)
                                }
                                else if(parseInt(e.currentTarget.value) < 1){
                                    setCurrentPagePending(1)
                                }
                                else{
                                    setCurrentPagePending(parseInt(e.currentTarget.value))
                                }
                            }
                        }
                    }
                    />
                    <input type="text" className=" bg-transparent w-10 text-center outline-none ring-0" value={`/  ${json1?.last_page}` } readOnly/>

                    <PaginationItem>
                        <PaginationNext className=" cursor-pointer" onClick={()=>
                            {
                                if(json1 && json1.page < json1.last_page && json1.last_page > 1){
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
                            success.map((leave) =>
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
                                    leave.status == 'approved' ? 
                                    <TableCell className=" flex flex-row">
                                        <CheckIcon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell> : 
                                    leave.status == "denied" ?
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
                                    if(json2 && json2.page > 1){
                                        setCurrentPageSuccess(currentPageSuccess - 1)
                                    }}}
                            />
                    </PaginationItem>
                    <Input type="number" className=" w-10" value={currentPageSuccess} onChange={(e)=>
                        {
                            if (json2){
                                if (e.currentTarget.value === ""){
                                    setCurrentPageSuccess(1)
                                }
                                else if(parseInt(e.currentTarget.value) > json2.last_page){
                                    setCurrentPageSuccess(json2.last_page)
                                }
                                else if(parseInt(e.currentTarget.value) < 1){
                                    setCurrentPageSuccess(1)
                                }
                                else{
                                    setCurrentPageSuccess(parseInt(e.currentTarget.value))
                                }
                            }
                        }
                    }
                    />
                    <input type="text" className=" bg-transparent w-10 text-center outline-none ring-0" value={`/  ${json2?.last_page}` } readOnly/>

                    <PaginationItem>
                        <PaginationNext className=" cursor-pointer" onClick={()=>
                            {
                                if(json2 && json2.page < json2.last_page && json2.last_page > 1){
                                    setCurrentPageSuccess(currentPageSuccess + 1)
                                }}
                            }
                            
                            />
                    </PaginationItem>
                </PaginationContent>
                </Pagination>
            </div>
            <div>
                <h1 className="font-bold text-2xl">
                    Download Leaves Data
                </h1>
            </div>
            <div className="self-center">
                {downloadUrl ? 
                    <a href={downloadUrl} download="leaves.csv">
                        <Button>
                            Download
                        </Button>
                    </a> : null
                }
            </div>
        </main>
    )
}