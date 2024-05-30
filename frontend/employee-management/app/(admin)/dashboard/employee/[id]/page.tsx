"use client";

import { Button } from "@/components/ui/button";
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectTrigger,
    SelectValue,
  } from "@/components/ui/select"
  import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
  } from "@/components/ui/card"
import { format } from "date-fns"
import { cn } from "@/lib/utils"
import { Label } from "@/components/ui/label";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import GetOneEmployee from "@/lib/GetOneEmployee";
import { Attendance, UserJson } from "@/interface";
import { useRouter } from "next/navigation";
import GetMyAttendances from "@/lib/GetMyAttendances";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
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
    Table,
    TableCaption,
    TableBody, 
    TableCell, 
    TableHead, 
    TableHeader, 
    TableRow 
} from "@/components/ui/table";
import { TextAlignBottomIcon , TextAlignTopIcon, CalendarIcon} from '@radix-ui/react-icons'
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
  } from "@/components/ui/popover"
import { Calendar } from "@/components/ui/calendar";
dayjs.extend(utc);


export default function Page({params} : {params: {id: string}}) {
    const [loading, setLoading] = useState<boolean>(false)
    const [employee, setEmployee] = useState<UserJson>()
    const router = useRouter();
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
            GetOneEmployee(session.user.token, params.id).then((data) => {
                setEmployee(data);
            }).catch((error) => {
                console.log(error);
            });
            if(date === undefined){
        
                if(selectedOption === "month"){
                    GetMyAttendances(params.id, session.user.token).then((res) => {
                        setData(sortItem(res.filter((att : Attendance) => {
                            return dayjs(att.date).local().format('MM/YYYY') === dayjs(date).local().format('MM/YYYY')
                        })));
                    });
                }
                else if (selectedOption === "year"){
                    GetMyAttendances(params.id, session.user.token).then((res) => {
                        setData(sortItem(res.filter(
                            (att : Attendance) => {
                                return dayjs(att.date).local().format('YYYY') === dayjs(date).local().format('YYYY')
                            }
                        )));
                    });
                }
                else {
                    GetMyAttendances(params.id, session.user.token).then((res) => {
                        setData(sortItem(res));
                    });
                }
            }
            else{
                GetMyAttendances(params.id, session.user.token).then((res) => {
                    setData(res.filter((att : Attendance) => {
                        return dayjs(att.date).local().format('DD/MM/YYYY') === dayjs(date).local().format('DD/MM/YYYY')
                        }));
                });
            }
        }, [selectedOption , date , sort]
    )

    return (
        <div className="flex flex-col md:space-y-7 space-y-5 px-[5%]  py-[5%] w-screen md:w-[80%] 2xl:w-[60%] gap-[5%]">
            <Button 
            onClick={() => {
                router.push("/dashboard/employee")
            }
            }
            className=" w-20 flex justify-center items-center">
                Back
            </Button>
            <h1 className="text-2xl font-bold">Employee ID : {params.id}</h1>
            <div className="flex flex-row space-x-3 justify-between w-full">
                <div className=" w-2/12 flex flex-col space-y-3">
                    <Label htmlFor="title_eng">Title(ENG)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.title_en}
                    </div>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="name_eng">Name(ENG)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.first_name_en}
                    </div>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="surname_eng">Surname(ENG)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.last_name_en}
                    </div>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full">
                <div className=" w-2/12 flex flex-col space-y-3">
                    <Label htmlFor="title_th">Title(TH)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.title_th}
                    </div>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="name_th">Name(TH)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.first_name_th}
                    </div>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="surname_th">Surname(TH)</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.last_name_th}
                    </div>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full">
            
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="date_of_birth">Date of birth</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.date_of_birth}
                    </div>
                </div>
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="surname_th">Gender</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.gender}
                    </div>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full"> 
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="department">Department</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.department}
                    </div>
                </div>
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="phone">Phone</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.phone}
                    </div>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full"> 
                <div className=" w-full flex flex-col space-y-3 ">
                    <Label htmlFor="email">Email</Label>
                    <div className="w-full flex items-center h-10 border border-gray-300 rounded-md px-3">
                        {employee?.email}
                    </div>
                </div>
            </div>
            <div className=" flex flex-col space-y-7">
                <Card className="">
                    <CardHeader className="">
                        <CardTitle>
                            History
                        </CardTitle>
                        <CardDescription>
                            history of employee
                        </CardDescription> 
                    </CardHeader>
                    <CardContent className=" flex flex-row space-x-7">
                        <div className=" border w-1/2 h-40 p-5 rounded-md flex flex-col space-y-5">
                            <h1 className=" text-xl">
                                Attendance
                            </h1>
                            <h1 className=" text-2xl">
                                {data.length}
                            </h1>

                        </div>
                        <div className=" border w-1/2 h-40 p-5 rounded-md flex flex-col space-y-5">
                            <h1 className=" text-xl">
                                Leave
                            </h1>
                            <h1 className=" text-2xl">
                                {
                                    data.filter((att : Attendance) => {
                                        return att.leave_id !== -1
                                    }).length

                                }

                                
                            </h1>
                        </div>
                        
                    </CardContent>
                </Card>
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
            </div>
            
            
        </div>
    );
}