"use client"

import { BellIcon  } from "@radix-ui/react-icons"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import GetNotification from "@/lib/GetNotification"
import { useEffect, useState } from "react"
import { useSession } from "next-auth/react"
import { Notification } from "@/interface"
import { CheckIcon , Cross1Icon , CalendarIcon } from "@radix-ui/react-icons"
import {
  Card,
  CardContent,
} from "@/components/ui/card"
import MakeReadNotification from "@/lib/MakeReadNotification"

export  function MenuBar() {
  const { data: session } = useSession()
  const [notifications, setNotifications] = useState<Notification[] | []>([])

  useEffect(() => {
    if(session){
      GetNotification(session?.user.token, session?.user.employee_id, session?.user.role).then((data) => {
        if(session?.user.role === "admin"){
          setNotifications(data)
          setNotifications(data.filter((notification : Notification) => notification.title === "Leave Request" && notification.read === false))
        }
        else{
          setNotifications(data)
          setNotifications(data.filter((notification : Notification) => notification.title !== "Leave Request" && notification.read === false))
        }
      })
    }
  })

  const handleNotification = (nid : Number) => {
    if(session){
      MakeReadNotification(session?.user.token, nid)
    }
   
  }
  
  return (
    <Popover>
      <PopoverTrigger asChild>
        <button className=" relative">
          <BellIcon className=" mr-2 h-6 w-6 " />
          {notifications.length > 0 && (
            <span className="absolute text-white text-sm font-semibold -top-1 right-.5 h-5 w-5  text-center rounded-full bg-red-500">
              {notifications.length}
            </span>
          )}
        </button>
      </PopoverTrigger>
      <PopoverContent className="w-60 mr-20">
        <div className="grid gap-4">
          <div className="space-y-2">
            <h4 className="font-medium leading-none">Notifications</h4>
            <hr />
            <div className=" h-40 space-y-2 overflow-y-scroll scrollbar-hide">
            {
              notifications.length > 0 
              ? 
              (
                session?.user.role === "admin" ?
                (
                  notifications.filter(
                  (notification : Notification) => notification.title === "Leave Request"
                ).map((notification : Notification) => (
                <Card className=" relative">
                  <CardContent>
                    <div className=" flex flex-row w-full space-x-3 justify-between pt-7">
                      <CalendarIcon className="h-6 w-6 " />
                      <div className=" flex flex-col">
                        <p>{notification.title}</p>
                        <p className=" text-xs">
                          {notification.message}
                        </p>
                      </div>
                    </div>
                  </CardContent>
                  <button onClick={()=>handleNotification(notification.id)} className="absolute right-2 top-2">
                    <Cross1Icon className="h-4 w-4" />
                  </button>
                </Card>)
              )) 
              :
              (
                notifications.filter(
                  (notification : Notification) => notification.title !== "Leave Request"
                ).map((notification : Notification) => (
                <Card className=" relative">
                  <CardContent>
                    <div className=" flex flex-row w-full space-x-3 justify-between pt-7">
                      {
                        notification.title === "Leave Approved" ? <CheckIcon className="h-6 w-6" /> 
                        : <Cross1Icon className="h-6 w-6 " />
                      }
                      
                      <div className=" flex flex-col">

                        <p>{notification.title}</p>
                        <p className=" text-xs">
                          {notification.message}
                        </p>
                      </div>
                    </div>
                  </CardContent>
                  <button onClick={()=>handleNotification(notification.id)} className="absolute right-2 top-2">
                    <Cross1Icon className="h-4 w-4" />
                  </button>
                </Card>)
              )
              )
              )
              : 
              <div className=" w-full h-40 flex justify-center items-center">
                <p className="text-sm">No notifications</p>
              </div>
            }
            </div>
          </div>
        </div>
      </PopoverContent>
    </Popover>
  )
}
