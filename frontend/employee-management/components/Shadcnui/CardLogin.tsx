"use client";
import { signIn, } from "next-auth/react";
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


export function CardLogin({setPass , setEmail , onSubmit} : {setPass : any , setEmail : any , onSubmit : any}) {
    
  
  return (
    <Card className="w-[550px]">
      <CardHeader>
        <CardTitle>Login</CardTitle>
        <CardDescription>Enter your email and password below to login your account</CardDescription>
      </CardHeader>
      <CardContent>
        <form>
          <div className="grid w-full items-center gap-4">
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="email">Email</Label>
              <Input onChange={(e)=>{setEmail(e.currentTarget.value)}} type="email" id="email" placeholder="john.doe@gmail.com" />
            </div>
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="password">Password</Label>
              <Input onChange={(e)=>{setPass(e.currentTarget.value)}} type="password" id="password" placeholder="Your password" />
            </div>

          </div>
        </form>
      </CardContent>
      <CardFooter className="flex justify-between">
        <Button onClick={()=>onSubmit} className=" w-full text-center flex justify-center items-center"> Login</Button>
      </CardFooter>
      <CardFooter className="flex justify-between">
        <CardDescription className=" flex text-center px-16">By clicking continue, you agree to our Terms of Service and Privacy Policy.</CardDescription>
      </CardFooter>
    </Card>
  )
}
