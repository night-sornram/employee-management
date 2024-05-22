"use client"

import { Button } from "@/components/ui/button";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
  } from "@/components/ui/select"
import { Label } from "@/components/ui/label";
import { useSession } from "next-auth/react";
import { useState } from "react";
import CreateEmployee from "@/lib/CreateEmployee";
import { set } from "date-fns";
import { ReloadIcon } from "@radix-ui/react-icons";

export default function LeaveRequestPage () {
    const { data: session } = useSession()
    const [id, setId] = useState<string>('')
    const [title_eng, setTitleEng] = useState<string>('')
    const [name_eng, setNameEng] = useState<string>('')
    const [surname_eng, setSurnameEng] = useState<string>('')
    const [title_th, setTitleTh] = useState<string>('')
    const [name_th, setNameTh] = useState<string>('')
    const [surname_th, setSurnameTh] = useState<string>('')
    const [dateOfBirth, setDateOfBirth] = useState<string>('')
    const [gender , setGender] = useState<string>('')
    const [department, setDepartment] = useState<string>('')
    const [phone, setPhone] = useState<string>('')
    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    const [loading, setLoading] = useState<boolean>(false)

    const onSubmit = async () => {
        if (id === '' || title_eng === '' || name_eng === '' || surname_eng === '' || title_th === '' || name_th === '' || surname_th === '' || dateOfBirth === '' || gender === '' || department === '' || phone === '' || email === '' || password === '') {
            alert('Please fill all fields')
        }
        else{
            if (session) {
                let data : object = {
                    employee_id: id,
                    title_en: title_eng,
                    first_name_en: name_eng,
                    last_name_en: surname_eng,
                    title_th: title_th,
                    first_name_th: name_th,
                    last_name_th: surname_th,
                    date_of_birth: dateOfBirth,
                    gender : gender,
                    department: department,
                    phone: phone,
                    role : "user",
                    email: email,
                    password: password
                }
                try{
                    setLoading(true)
                    await CreateEmployee(session.user.token, data)
                    alert('Employee created successfully')
                    setLoading(false)
                    window.location.reload()
                }
                catch (error) {
                    console.log(error)
                }
            }
        }
        
    }
    
    return (
        <div className="flex flex-col px-[10%] py-[5%] w-[60%] gap-[5%]">
            <h1 className="text-2xl font-bold">Create Employee</h1>
            <div className=" w-12/12 flex flex-col space-y-3">
                <Label htmlFor="id">Employee ID</Label>
                <input placeholder="A123" onChange={(e)=>{setId(e.currentTarget.value)}} type="text" name="id" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full">
                <div className=" w-2/12 flex flex-col space-y-3">
                    <Label htmlFor="title_eng">Title(ENG)</Label>
                    <input placeholder="Mr." onChange={(e)=>setTitleEng(e.currentTarget.value)} type="text" name="title_eng" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="name_eng">Name(ENG)</Label>
                    <input placeholder="Somchai" onChange={(e)=>setNameEng(e.currentTarget.value)} type="text" name="name_eng" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="surname_eng">Surname(ENG)</Label>
                    <input placeholder="Jaide" onChange={(e)=>setSurnameEng(e.currentTarget.value)} type="text" name="surname_eng" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full">
                <div className=" w-2/12 flex flex-col space-y-3">
                    <Label htmlFor="title_th">Title(TH)</Label>
                    <input placeholder="นาย" onChange={(e)=>setTitleTh(e.currentTarget.value)} type="text" name="title_th" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="name_th">Name(TH)</Label>
                    <input placeholder="สมชาย" onChange={(e)=>setNameTh(e.currentTarget.value)} type="text" name="name_th" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
                <div className=" w-5/12 flex flex-col space-y-3">
                    <Label htmlFor="surname_th">Surname(TH)</Label>
                    <input placeholder="ใจดี" onChange={(e)=>setSurnameTh(e.currentTarget.value)} type="text" name="surname_th" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full">
            
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="date_of_birth">Date of birth</Label>
                    <input placeholder="1990-01-01" onChange={(e)=>setDateOfBirth(e.currentTarget.value)} type="text" name="date_of_birth" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="surname_th">Gender</Label>
                    <Select 
                    value={gender}
                    onValueChange={setGender} 
                    name="gender">
                        <SelectTrigger  className="w-full">
                            <SelectValue placeholder="Gender" />
                        </SelectTrigger>
                        <SelectContent>
                            <SelectItem value="male">Male</SelectItem>
                            <SelectItem value="female">Female</SelectItem>
                            <SelectItem value="other">Other</SelectItem>
                        </SelectContent>
                    </Select>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full"> 
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="department">Department</Label>
                    <input placeholder="IT" onChange={(e)=>setDepartment(e.currentTarget.value)} type="text" name="department" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="phone">Phone</Label>
                    <input placeholder="080-000-0000" onChange={(e)=>setPhone(e.currentTarget.value)} type="text" name="phone" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full"> 
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="email">Email</Label>
                    <input placeholder="somchai.j@gmail.com" onChange={(e)=>setEmail(e.currentTarget.value)} type="text" name="email" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
                <div className=" w-6/12 flex flex-col space-y-3">
                    <Label htmlFor="password">Password</Label>
                    <input placeholder="123456" onChange={(e)=>setPassword(e.currentTarget.value)} type="text" name="password" className="w-full h-10 border border-gray-300 rounded-md px-3"/>
                </div>
            </div>
            
            <div className="items-center w-full text-center ">
                {
                    loading ? 
                    (
                        <Button disabled className=" w-full flex justify-center">
                            <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />
                            Please wait
                        </Button>
                    )
                    :
                    (
                        <Button onClick={onSubmit} className=" w-full flex justify-center" >Submit</Button>
                    )
                }
            </div>
        </div>
    );
}