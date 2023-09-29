'use client'
import EnrollmentHoc from "@/app/auth/auth";
import CardAuth from "@/components/organisms/CardAuth";
import LogoTitle from "@/components/atoms/SearchComponents/titleSearch";

import { WInput } from "@/components";

export default function LoginPage() {

  return (
    <EnrollmentHoc>
        <CardAuth title="testing-title-login-user">
            <span>Email</span>
            <WInput typeColor="primary" placeholder="tesing usario" fullWidth/>
        </CardAuth>
    </EnrollmentHoc>
  );
}