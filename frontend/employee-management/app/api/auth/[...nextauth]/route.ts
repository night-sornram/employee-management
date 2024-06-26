import NextAuth from "next-auth/next";
import { AuthOptions } from "next-auth";
import  CredentialsProvider from "next-auth/providers/credentials";
import UserLogin from "@/lib/UserLogin";
import getUserProfile from "@/lib/GetUserProfile";


export const authOptions:AuthOptions = {
    providers: [
        CredentialsProvider({
          name: "Credentials",
          credentials: {
            id: { label: "id", type: "string", placeholder: "id" },
            password: { label: "Password", type: "password" }
          },
          async authorize(credentials, req) {

            if(!credentials) return null

            const user = await UserLogin(credentials.id, credentials.password)

            if (user) {
              return user
            } else {
              return null
            }
          }
        })
      ],
    secret: process.env.NEXTAUTH_SECRET,
    session: { strategy : "jwt" },
    callbacks: {
        async jwt({token, user}) {
          return {...token, ...user }
        },
        async session({session, token,user}) {
          if(token.token){
            
            const res = await getUserProfile(token.token as string)
            session.user.email = res.email
            session.user.employee_id = res.employee_id
            session.user.name = res.name
            session.user.role = res.role
            token.role = res.role
            session.user.token = token.token as string

          }

          return session
        },
        async redirect({ url, baseUrl }) {
          return url.startsWith(baseUrl) ? url : baseUrl
        }
      }
}

const handler = NextAuth(authOptions)
export {handler as GET, handler as POST , handler as PUT}