import { withAuth } from "next-auth/middleware";
import { NextResponse } from "next/server";

export default withAuth(
  function middleware(req) {
    if (
     ( (req.nextUrl.pathname === "/dashboard/approve-leave" || req.nextUrl.pathname === "/dashboard/create-employee" || req.nextUrl.pathname === "/dashboard/all-attendance-history" ) &&
      req.nextauth.token?.role !== "admin")
    ) {
      return NextResponse.redirect(new URL('/', req.url))
    }
  },
  {
    callbacks: {
      authorized: (params) => {
        let { token } = params;
        return !!token;
      },
    },
  }
);

export const config = {
    matcher: ['/attendance' , '/attendance/:path*' , '/leave-request' , '/leave-request/:path*' , '/setting' , '/setting/:path*' , "/dashboard" , "/dashboard/:path*" ],
};

