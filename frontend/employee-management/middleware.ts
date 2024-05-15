export { default } from 'next-auth/middleware'

export const config = {
    matcher: ['/attendance' , '/attendance/:path*' , '/leave-request' , '/leave-request/:path*' , '/setting' , '/setting/:path*' ],
};

