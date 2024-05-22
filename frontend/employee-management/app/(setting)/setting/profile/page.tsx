"use client"

import { useSession } from "next-auth/react"
import { ReloadIcon } from "@radix-ui/react-icons"
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
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
  } from "@/components/ui/dialog"
import GetUserProfile from "@/lib/GetUserProfile"
import { UserJson } from "@/interface"
import { Skeleton } from "@/components/ui/skeleton"
import ChangeEmail from "@/lib/ChangeEmail"
import ChangePhone from "@/lib/ChangePhone"
import ChangePassword from "@/lib/ChangePassword"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"



export default  function Page(){
    const { data: session } = useSession()
    const [user, setUser] = useState<UserJson | null>(null)
    const [loading, setLoading] = useState(true)
    const [loadingPassword, setLoadingPassword] = useState(false)
    const [email, setEmail] = useState("")
    const [phone, setPhone] = useState("")
    const [currentPassword, setCurrentPassword] = useState("")
    const [newPassword, setNewPassword] = useState("")
    const [repeatPassword, setRepeatPassword] = useState("")

    useEffect(() => {
            if (session?.user.token) {
                 GetUserProfile(session.user.token).then((res) => {
                    setUser(res)
                    setLoading(false)
                    setEmail(res.email)
                    setPhone(res.phone)
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

    const handleChangeEmail = async () => {
        try {
            if (user) {
                await ChangeEmail(session.user.token, email, user.employee_id)
                window.location.reload()
            } else {
                throw new Error("User is null")
            }
        } catch (error) {
            alert("Failed to change email")
        }
    
    }

    const handleChangePhone = async () => {
        try {
            if (user) {
                await ChangePhone(session.user.token, phone, user.employee_id)
                window.location.reload()
            } else {
                throw new Error("User is null")
            }
        } catch (error) {
            alert("Failed to change phone")
        }
    
    }

    const handleChangePassword = async () => {
        if (newPassword === "" || repeatPassword === "" || currentPassword === "") {
            alert("Please fill all the field")
        }
        else{
            if(newPassword !== repeatPassword){
                alert("Password does not match")
            }
            else{
                try {
                    if (user) {
                        setLoadingPassword(true)
                        await ChangePassword(session.user.token, currentPassword, newPassword, user?.employee_id)
                        alert("Password has been changed")
                        setNewPassword("")
                        setRepeatPassword("")
                        setCurrentPassword("")
                    } else {
                        throw new Error("User is null")
                    }
                } catch (error) {
                    alert("Failed to change password")
                }
            }
        }
        setLoadingPassword(false)
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
            <div className=" flex flex-col space-y-5">
                <Card className="w-[50vw] ">
                    <CardContent>
                        {
                            loading ? 
                            (
                                <div className="flex pt-6 items-center space-x-4">
                                    <Skeleton className="h-20 w-20 rounded-full" />
                                    <div className="space-y-2">
                                        <Skeleton className="h-4 w-[250px]" />
                                        <Skeleton className="h-4 w-[200px]" />
                                    </div>
                                </div>
                            )
                            
                            
                            : (
                                <div className=" flex pt-6 flex-row items-center space-x-10 ">
                                    <Avatar className=" w-20 h-20">
                                        <AvatarImage src="https://github.com/shadcn.png" />
                                        <AvatarFallback>CN</AvatarFallback>
                                    </Avatar>
                                    <div className=" flex flex-col justify-start items-start space-y-2">
                                        <Label className=" text-xl font-bold" htmlFor="name">{user?.title_en + " " +  user?.first_name_en + "  " +   user?.last_name_en}</Label>
                                        <Label className=" text-lg " htmlFor="name">{user?.title_th + " " +  user?.first_name_th + "  " +   user?.last_name_th}</Label>
                                    </div>
                                </div>
                            )

                        }
                    </CardContent>
                </Card>
                <Card className="w-[50vw] p-4 ">
                    <div className=" flex flex-row justify-between items-center">
                        <div className=" flex flex-row space-x-5 items-center">            
                            <Label className=" text-lg font-bold" htmlFor="email">Email</Label>
                            {
                                loading ? 
                                (
                                    <Skeleton className="h-4 w-[200px]" />
                                )
                                : (
                                    <Label className="" htmlFor="email">{user?.email}</Label>   
                                )
                            }
                        </div>
                        <Dialog>
                            <DialogTrigger asChild>
                                <Button variant="outline">Edit Email</Button>
                            </DialogTrigger>
                            <DialogContent className="sm:max-w-[425px]">
                                <DialogHeader>
                                <DialogTitle>Edit Email</DialogTitle>
                                <DialogDescription>
                                    Update your email address
                                </DialogDescription>
                                </DialogHeader>
                                <div className="grid gap-4 py-4">
                                    <div className="grid grid-cols-4 items-center gap-4">
                                        <Input onChange={(e)=>{setEmail(e.currentTarget.value)}} id="username" value={email} className="col-span-4" />
                                    </div>
                                </div>
                                <DialogFooter>
                                <Button onClick={handleChangeEmail} type="submit">Save changes</Button>
                                </DialogFooter>
                            </DialogContent>
                        </Dialog>
                    </div>
                </Card>
                <Card className="w-[50vw] p-4 ">
                    <div className=" flex flex-row justify-between items-center">
                        <div className=" flex flex-row space-x-5 items-center">            
                            <Label className=" text-lg font-bold" htmlFor="email">Phone</Label>
                            {
                                loading ? 
                                (
                                    <Skeleton className="h-4 w-[200px]" />
                                )
                                : (
                                    <Label className="" htmlFor="email">{user?.phone}</Label>   
                                )
                            }
                        </div>
                        <Dialog>
                            <DialogTrigger asChild>
                                <Button variant="outline">Edit Phone</Button>
                            </DialogTrigger>
                            <DialogContent className="sm:max-w-[425px]">
                                <DialogHeader>
                                <DialogTitle>Edit Phone</DialogTitle>
                                <DialogDescription>
                                    Update your phone number
                                </DialogDescription>
                                </DialogHeader>
                                <div className="grid gap-4 py-4">
                                    <div className="grid grid-cols-4 items-center gap-4">
                                        <Input onChange={(e)=>{setPhone(e.currentTarget.value)}} id="username" value={phone} className="col-span-4" />
                                    </div>
                                </div>
                                <DialogFooter>
                                <Button onClick={handleChangePhone} type="submit">Save changes</Button>
                                </DialogFooter>
                            </DialogContent>
                        </Dialog>
                    </div>
                </Card>
                <Card className="w-[50vw]">
                    <CardHeader>
                        <CardTitle>Change Password</CardTitle>
                        <CardDescription>
                            Update your password
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <div className=" space-y-6">
                            <div className="grid w-full items-center gap-4">
                                <div className="flex flex-col space-y-1.5">
                                    <Label htmlFor="currentPassword">Current Password</Label>
                                    <Input onChange={(e)=>{setCurrentPassword(e.currentTarget.value)}} value={currentPassword} id="currentPassword" placeholder=" your current password " />
                                </div>
                            </div>
                            <div className="grid w-full items-center gap-4">
                                <div className="flex flex-col space-y-1.5">
                                    <Label htmlFor="currentPassword">New Password</Label>
                                    <Input onChange={(e)=>{setNewPassword(e.currentTarget.value)}} value={newPassword} id="currentPassword" placeholder=" new password " />
                                    <Input onChange={(e)=>{setRepeatPassword(e.currentTarget.value)}} value={repeatPassword} id="currentPassword" placeholder=" repeat new password " />
                                </div>
                            </div>
                            <div className="grid w-full items-center gap-4 ">
                                {
                                    loadingPassword ? 
                                    (
                                        <Button className=" flex justify-center " disabled>
                                            <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />
                                            Please wait
                                        </Button>
                                    )
                                    : 
                                    (
                                        <Button  className=" flex justify-center" onClick={handleChangePassword}>Save changes</Button>
                                    )
                                }
                            </div>
                            
                        </div>
                    </CardContent>
                    </Card>

            </div>
        </div>
    )
}