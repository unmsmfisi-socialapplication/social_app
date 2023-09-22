"use client";
import { useEffect, useState } from "react";
import { Box, Grid } from "@mui/material";

import RootEnrollment from "../Enrollment";
import { WButton } from "@/components";

export default function EnrollmentHoc({
  children,
}: {
  children: React.ReactNode;
}) {
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  return (
    <RootEnrollment>
      {isClient && (
        <Grid container spacing={2}>
          <Grid
            style={{
              paddingRight: "240px",
              color: "white",
              fontFamily: "inherit",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
            }}
            item
            xs={8}
          >
            <span style={{ fontSize: "80px" }}>STUDENT NETWORK</span>
          </Grid>
          <Grid item xs={4}>
            <Box>{children}</Box>
          </Grid>
        </Grid>
      )}
    </RootEnrollment>
  );
}
