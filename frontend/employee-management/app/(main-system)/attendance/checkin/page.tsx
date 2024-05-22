"use client"

import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input"
import GetUserProfile from "@/lib/GetUserProfile";
import { useState , useEffect } from "react";
import { useSession } from "next-auth/react";
import { Attendance, UserJson } from "@/interface";
import { Skeleton } from "@/components/ui/skeleton"
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert"
import { Cross1Icon , CheckIcon , CalendarIcon } from "@radix-ui/react-icons"
import GetTodayCheckIn from "@/lib/GetTodayCheckIn";
import Checkin from "@/lib/Checkin";
import { useToast } from "@/components/ui/use-toast"
import CreateNotification from "@/lib/CreateNotification";



export default function Page() {

    const { data: session } = useSession()
    const [user, setUser] = useState<UserJson | null>(null)
    const [loading, setLoading] = useState(true)
    const [time , setTime] = useState(new Date())
    const [data, setData] = useState<Attendance | null>(null)
    const { toast } = useToast()


    useEffect(() => {
            setTime(new Date())
            if (session?.user.token) {
                 GetUserProfile(session.user.token).then((res) => {
                    setUser(res)
                    setLoading(false)
                 })
                GetTodayCheckIn(session.user.employee_id, session.user.token).then((res) => {
                    setData(res)
                })
            }
            
    },[])

    if (!session) {
            return {
                redirect: {
                    destination: "/",
                    permanent: false,
                },
            }
        }

    const checkIn = async () => {
        try {
            await Checkin(session.user.token, session.user.employee_id)
            await CreateNotification(session.user.token, `Check In`,`You have checked-In on ${new Date().toLocaleDateString('th-TH')}`, false, session.user.employee_id);
            toast({
                title: "Check-In Success",
                description: "You have checked in successfully",
              })
            setTimeout(() => {
                window.location.reload()
            }
            , 1000)
        } catch (error) {
            toast({
                title: "Check-In Failed",
                variant: "destructive",
                description: "Please try again",
              })
        }
    }
    
    return(
        <div className="flex flex-col px-[10%] py-[5%] w-[60%] gap-[5%]">
            {
                loading ? 
                (
                    <Skeleton className="w-full h-20"/>

                )
                :
                (
                    
                    data === null ?(
                        <Alert className=" w-full h-20" variant="destructive">
                            <Cross1Icon className="h-4 w-4" />
                            <AlertTitle>
                                NOT CHECKED IN
                            </AlertTitle>
                            <AlertDescription>
                                Please check in before 10:00
                            </AlertDescription>
                        </Alert> 
                    )
                    :
                    (
                        data.leave_id !== -1 ? (
                            
                            <Alert className=" w-full h-20" variant="default">
                                <CalendarIcon className="h-4 w-4" />
                                <AlertTitle>
                                    LEAVE 
                                </AlertTitle>
                                <AlertDescription>
                                    Today have leave request
                                </AlertDescription>
                            </Alert>
                        )
                        :
                        (
                            <Alert className=" w-full h-20" variant="default">
                                <CheckIcon className="h-4 w-4" />
                                <AlertTitle>
                                    ALREADY CHECKED IN
                                </AlertTitle>
                                <AlertDescription>
                                    You have checked in at {new Date(data.check_in).toLocaleTimeString("th-TH")}
                                </AlertDescription>
                            </Alert> 
                            
                        )
                    )          
                )
                
            }
            <h1 className="text-2xl font-bold">Check In</h1>
            <div className="flex flex-col space-y-3 justify-between w-full">
                <Label htmlFor="reason">Name</Label>
                {
                    loading  ?
                    (
                        <Skeleton className="w-full h-10" />
                    )
                    : 
                    (
                        <Input disabled type="string" placeholder={ user?.title_en + " " + user?.first_name_en + " " + user?.last_name_en} />
                    )
                }
        </div>
        <div className=" flex flex-col space-y-3">
            <Label htmlFor="reason">Time</Label>
            {
                loading  ?
                (
                    <Skeleton className="w-full h-10" />
                )
                : 
                ( 
                    data === null ?(
                        <Input disabled type="string" placeholder={time.toLocaleString()} /> 
                    )
                    :
                    (
                        data.leave_id !== -1 ? (
                            <Input disabled type="string" placeholder="Leave" />
                        )
                        :
                        (
                            <Input disabled type="string" placeholder={new Date(data.check_in).toLocaleTimeString("th-TH")} />

                        )
                    )       
                    
                )
            }
        </div>
        <div className="items-center w-full text-center ">
            {
                loading ?
                (
                    <Skeleton className=" w-full h-10" />
                ) 
                :
                (
                    data === null ?(
                        <Button onClick={checkIn} className=" w-full flex justify-center" >Check-In</Button>
                    )
                    :
                    (
                        data.leave_id === -1 ? (
                            <Button disabled className=" w-full flex justify-center" >Already Checked-In</Button>
                        )
                        :
                        (
                            <Button disabled className=" w-full flex justify-center" >Check-In</Button>
                        )
                    )
                )
            }
        </div>
    </div>
    )
}