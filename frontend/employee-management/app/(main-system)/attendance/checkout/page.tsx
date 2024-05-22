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
import { Cross1Icon , CheckIcon , ExitIcon , CalendarIcon } from "@radix-ui/react-icons"
import GetTodayCheckIn from "@/lib/GetTodayCheckIn";
import { useToast } from "@/components/ui/use-toast"
import Checkout from "@/lib/Checkout";


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

    const handleCheckOut = async () => {
        try {
            await Checkout(session.user.token, data?.id as Number)
            toast({
                title: "Check-Out Success",
                description: "You have checked out successfully",
              })
              setTimeout(() => {
                window.location.reload()
            }, 1000)
        } catch (error) {
            toast({
                title: "Check-Out Failed",
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
                    <Skeleton className=" w-full h-20" />
                ) 
                :
                (
                    data === null ? 
                    (
                        <Alert className=" w-full h-20" variant="destructive">
                            <Cross1Icon className="h-4 w-4" />
                            <AlertTitle>
                                NOT CHECKED IN
                            </AlertTitle>
                            <AlertDescription>
                                Please check in before check out
                            </AlertDescription>
                        </Alert> 
                    )
                    : 
                    (
                        data.leave_id !== -1   ?
                        (
                            <Alert className=" w-full h-20" variant="default">
                                <CalendarIcon className="h-4 w-4" />
                                <AlertTitle>
                                    Leave
                                </AlertTitle>
                                <AlertDescription>
                                    Today you have leave
                                </AlertDescription>
                            </Alert> 
                        )
                        :
                        ( 
                            data.check_out === "0001-01-01T07:00:00+07:00"?
                            (
                            <Alert className=" w-full h-20" variant="default">
                                <ExitIcon className="h-4 w-4" />
                                <AlertTitle>
                                    CHECK OUT
                                </AlertTitle>
                                <AlertDescription>
                                    Please check out before 18:00 
                                </AlertDescription>
                            </Alert> 
                            )
                            :
                            (
                                <Alert className=" w-full h-20" variant="default">
                                    <CheckIcon className="h-4 w-4" />
                                    <AlertTitle>
                                        ALREADY CHECKED OUT
                                    </AlertTitle>
                                    <AlertDescription>
                                        You have checked out at {new Date(data.check_out).toLocaleTimeString("th-TH")}
                                    </AlertDescription>
                                </Alert> 

                            )
                        
                        )
                    )
                )
            }
            <h1 className="text-2xl font-bold">Check Out</h1>
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
                        data === null ?
                        (
                            <Input disabled type="string" placeholder={time.toUTCString()} />
                        )
                        :
                        (
                            data.leave_id !== -1 ?
                            (
                                <Input disabled type="string" placeholder="Leave" />
                            )
                            :
                            (
                                data.check_out === "0001-01-01T07:00:00+07:00" ?
                                (
                                    <Input disabled type="string" placeholder={time.toUTCString()} />
                                )
                                :
                                (
                                    <Input disabled type="string" placeholder={new Date(data.check_out).toLocaleTimeString("th-TH")} />
                                )
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
                   data === null ?
                    (
                        <Button disabled className=" w-full flex justify-center" >Check-Out</Button>
                    )
                    :
                    (
                       data.leave_id !== -1 ?
                        (
                            <Button disabled className=" w-full flex justify-center" >Check-Out</Button> 
                        )
                        :
                        (
                            data.check_out === "0001-01-01T07:00:00+07:00"
                            ?   
                            (
                                <Button onClick={handleCheckOut}  className=" w-full flex justify-center" >Check-Out</Button>
                            )
                            :
                            (
                                <Button disabled className=" w-full flex justify-center" >Already Check-Out</Button>
                            )
                        )
                        
                    )
                )
            }
        </div>
    </div>
    )
}