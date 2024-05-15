"use client"

import { useSession } from "next-auth/react"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
  } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"
import { useEffect , useState} from "react"
import CircularProgress from '@mui/material/CircularProgress';
import GetUserProfile from "@/lib/GetUserProfile"
import { UserJson } from "@/interface"

export default  function Page(){
    const { data: session } = useSession()
    const [user, setUser] = useState<UserJson | null>(null)
    const [loading, setLoading] = useState(true)

    useEffect(() => {
            if (session?.user.token) {
                 GetUserProfile(session.user.token).then((res) => {
                    setUser(res)
                    setLoading(false)
                 })
            }
    })

    if (!session) {
            return {
                redirect: {
                    destination: "/",
                    permanent: false,
                },
            }
        }

    
    return(
        <div className=" w-[50vw] px-5 space-y-5">
            <div className=" flex flex-col">
                <h1 className=" text-lg font-medium">
                    Profile
                </h1>
                <div className=" text-gray-400 dark:text-gray-400">
                    Manage your account settings
                </div>
            </div>
            <hr />
            <div>
                <Card className="w-[50vw] ">
                    <CardContent>
                        {
                            loading ? 
                            (
                                <div className=" flex pt-6 justify-center items-center">
                                    <CircularProgress />
                                </div>
                            )
                            
                            
                            : (
                                <div className=" flex pt-6 flex-row items-center space-x-10 ">
                                    <div className=" w-[140px] h-[140px] bg-red-500 rounded-full">
                                    </div>
                                    <div className=" flex flex-col justify-start items-start space-y-3">
                                        <Label className=" text-base font-medium text-gray-400" htmlFor="name">{user?.employee_id}</Label>
                                        <Label className=" text-xl font-bold" htmlFor="name">{user?.title_en + " " +  user?.first_name_en + "  " +   user?.last_name_en}</Label>
                                        <Label className=" text-lg " htmlFor="name">{user?.title_th + " " +  user?.first_name_th + "  " +   user?.last_name_th}</Label>
                                    </div>
                                </div>
                            )

                        }
                        
                        
                    </CardContent>

                </Card>
            </div>


        </div>
    )
}