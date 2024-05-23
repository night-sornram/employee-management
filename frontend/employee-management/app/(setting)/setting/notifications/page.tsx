"use client"

import { RootState, useAppSelector } from '@/store/store'
import { useDispatch } from "react-redux"
import { AppDispatch } from "@/store/store"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
  } from "@/components/ui/card"
  import { BellIcon , EnterIcon , CalendarIcon } from '@radix-ui/react-icons'
  import { useState } from "react"
  import { updateNotification } from '@/store/slices/appSlice'
import { Button } from '@/components/ui/button'

export default function Page(){
    const dispatch = useDispatch<AppDispatch>()
    const data = useAppSelector((state: RootState) => state.appReducer)
    const [notification, setNotification] = useState(data.notification)
    const [email, setEmail] = useState(data.email)
    return(
        <div className=" w-[50vw] px-5 space-y-5 ">
            <div className=" flex flex-col">
                <h1 className=" text-lg font-medium">
                    Notification
                </h1>
                <p className=" text-gray-400">
                    Manage your notifications settings
                </p>
            </div>
            <hr />
            <Card >
                <CardHeader>
                    <CardTitle>Notifications</CardTitle>
                    <CardDescription>Choose what you want to be notified about.</CardDescription>
                </CardHeader>
                <CardContent className=' flex flex-col space-y-2'>
                    <div
                    onClick={() => setNotification("all")}
                    className={` ${notification=== "all" ? "bg-gray-100 " : ""} w-full h-16 cursor-pointer   rounded-md p-5 flex flex-row items-center space-x-3`}> 
                        <BellIcon className=' w-7 h-7'/>
                        <div className=' flex flex-col'>
                            <h1 className=' font-medium'>Everything</h1>
                            <p className=' text-gray-400'>
                                Receive all notifications
                            </p>
                        </div>
                    </div>
                    <div 
                    onClick={() => setNotification("check")}
                    className={` ${notification  === "check" ? "bg-gray-100 " : ""} w-full h-16 cursor-pointer   rounded-md p-5 flex flex-row items-center space-x-3`}> 
                        <EnterIcon className=' w-7 h-7'/>
                        <div className=' flex flex-col'>
                            <h1 className=' font-medium'>Check-In Check-Out</h1>
                            <p className=' text-gray-400'>
                                Receive notifications for check-in and check-out
                            </p>
                        </div>
                    </div>
                    <div 
                    onClick={() => setNotification("leave")}
                    className={` ${notification  === "leave" ? "bg-gray-100 " : ""} w-full h-16 cursor-pointer   rounded-md p-5 flex flex-row items-center space-x-3`}> 
                        <CalendarIcon className=' w-7 h-7'/>
                        <div className=' flex flex-col'>
                            <h1 className=' font-medium'>Leave Approve</h1>
                            <p className=' text-gray-400'>
                                Receive notifications for leave approve
                            </p>
                        </div>
                    </div>
                    <Button className=' flex justify-center'
                    onClick={() => dispatch(updateNotification({
                        notification : notification,
                        email : email
                    }))}>
                        Save
                    </Button>
                </CardContent>
                
                
            </Card>

        </div>
    )
}