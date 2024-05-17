"use client"

import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input"
import GetUserProfile from "@/lib/GetUserProfile";
import { useState , useEffect } from "react";
import { useSession } from "next-auth/react";
import { UserJson } from "@/interface";
import { Skeleton } from "@/components/ui/skeleton"
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert"
import { Cross1Icon , CheckIcon , ExitIcon , CalendarIcon } from "@radix-ui/react-icons"


export default function Page() {

    const { data: session } = useSession()
    const [user, setUser] = useState<UserJson | null>(null)
    const [loading, setLoading] = useState(true)
    const [time , setTime] = useState(new Date())
    const [mock, setMock] = useState("leave")


    useEffect(() => {
            setTime(new Date())
            if (session?.user.token) {
                 GetUserProfile(session.user.token).then((res) => {
                    setUser(res)
                    setLoading(false)
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

    return(
        <div className="flex flex-col px-[10%] py-[5%] w-[60%] gap-[5%]">
            {
                loading ?
                (
                    <Skeleton className=" w-full h-20" />
                ) 
                :
                (
                    mock === "not checked in" ? 
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
                        mock === "not checked out" ?
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
                            mock === "checked out" ?
                            (
                                <Alert className=" w-full h-20" variant="default">
                                    <CheckIcon className="h-4 w-4" />
                                    <AlertTitle>
                                        ALREADY CHECKED OUT
                                    </AlertTitle>
                                    <AlertDescription>
                                        You have checked out at 17:30
                                    </AlertDescription>
                                </Alert> 
                            )
                            :
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
                        mock === "not checked in" ?
                        (
                            <Input disabled type="string" placeholder={time.toUTCString()} />
                        )
                        :
                        (
                            mock === "not checked out" ?
                            (
                                <Input disabled type="string" placeholder={time.toUTCString()} />
                            )
                            :
                            (
                                mock === "checked out" ?
                                (
                                    <Input disabled type="string" placeholder="17:30" />
                                )
                                :
                                (
                                    <Input disabled type="string" placeholder="Leave" />
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
                    mock === "not checked in" ?
                    (
                        <Button disabled className=" w-full flex justify-center" >Check-Out</Button>
                    )
                    :
                    (
                        mock === "not checked out" ?
                        (
                            <Button  className=" w-full flex justify-center" >Check-Out</Button>
                        )
                        :
                        (
                            mock === "checked out" 
                            ?   
                            (
                                <Button disabled className=" w-full flex justify-center" >Already Check-Out</Button>
                            )
                            :
                            (
                                <Button disabled className=" w-full flex justify-center" >Check-Out</Button> 
                            )
                        )
                        
                    )
                )
            }
        </div>
    </div>
    )
}