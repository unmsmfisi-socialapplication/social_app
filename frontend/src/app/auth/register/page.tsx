"use client";
import EnrollmentHoc from "@/app/auth/auth";
import CardAuth from "@/components/organisms/CardAuth";

import { WInput } from "@/components";
import { WButton } from "@/components";

import VisibilityOffOutlinedIcon from "@mui/icons-material/VisibilityOffOutlined";

export default function RegisterPage() {
  return (
    <EnrollmentHoc>
      <CardAuth
        typeColor="secondary"
        title="SIGN-UP"
        size="large"
        variant="outlined"
        style={{ fontWeight: "900" }}
      >
        <div style={{ margin: "15px 0" }}>
          <span>Your Full Name</span>
          <WInput
            typeColor="primary"
            placeholder="Enter your name"
            size="small"
            variant="outlined"
            fullWidth
            type="text"
          />
        </div>
        <div style={{ margin: "15px 0" }}>
          <span>Email</span>
          <WInput
            typeColor="primary"
            placeholder="Enter your email"
            size="small"
            variant="outlined"
            fullWidth
            type="text"
          />
        </div>
        <div style={{ margin: "15px 0" }}>
          <span>Password</span>
          <WInput
            typeColor="primary"
            icon={<VisibilityOffOutlinedIcon />}
            placeholder="Enter your password"
            size="small"
            variant="outlined"
            fullWidth
            type="password"
          />
        </div>
        <div style={{ margin: "15px 0" }}>
          <span>Confirm Password</span>
          <WInput
            typeColor="primary"
            icon={<VisibilityOffOutlinedIcon />}
            placeholder="Enter your password"
            size="small"
            variant="outlined"
            fullWidth
            type="password"
          />
        </div>

        <WButton typeColor="secondary" text="Sign in" size="large" />
      </CardAuth>
    </EnrollmentHoc>
  );
}
