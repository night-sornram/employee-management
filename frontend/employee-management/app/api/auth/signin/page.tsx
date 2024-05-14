"use client";
import { signIn, useSession } from "next-auth/react";
import { redirect } from 'next/navigation'
import { useState } from "react";
import * as React from "react"

import { Button } from "@/components/ui/button"
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



const LoginPage = () => {
  
  const [email, setEmail] = useState("");
  const [pass, setPass] = useState("");
  const { data:session } = useSession();
  if(!session){
    const onSubmit = async (e : React.FormEvent<HTMLFormElement>) => {
        e.preventDefault(); 
        await signIn("credentials", {
          email: email, 
          password: pass,
          redirect: true,
          callbackUrl: "/",
    });
  };
  return (
    <main className="flex flex-row w-screen h-screen">
        <div className=" w-1/2 h-full bg-zinc-800">

        </div>
        <div className=" w-1/2 h-full flex bg-white justify-center items-center">
            <Card className="w-[550px]">
                <CardHeader>
                    <CardTitle>Login</CardTitle>
                    <CardDescription>Enter your email and password below to login your account</CardDescription>
                </CardHeader>
                <CardContent>
                    <form onSubmit={onSubmit}>
                      <div className="grid w-full items-center gap-4">
                          <div className="flex flex-col space-y-1.5">
                          <Label htmlFor="email">Email</Label>
                          <Input onChange={(e)=>{setEmail(e.currentTarget.value)}} type="email" id="email" placeholder="john.doe@gmail.com" />
                          </div>
                          <div className="flex flex-col space-y-1.5">
                          <Label htmlFor="password">Password</Label>
                          <Input onChange={(e)=>{setPass(e.currentTarget.value)}} type="password" id="password" placeholder="Your password" />
                          </div>
                          <Button type="submit" className=" w-full text-center flex justify-center items-center"> Login</Button>
                      </div>
                    </form>
                </CardContent>
                <CardFooter className="flex justify-between">
                    <CardDescription className=" flex text-center px-16">By clicking continue, you agree to our Terms of Service and Privacy Policy.</CardDescription>
                </CardFooter>
            </Card>
        </div>

      
     
    </main>
  );
}else {
  
  redirect('/')
}
};

export default LoginPage;