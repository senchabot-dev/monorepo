// src/pages/_app.tsx
//import { withTRPC } from "@trpc/next";
import { SessionProvider } from "next-auth/react";
//import superjson from "superjson";
import type { AppType } from "next/app";
//import type { AppRouter } from "../server/router";
import type { Session } from "next-auth";
import "../styles/globals.css";
import "./App.css";
import { trpc } from "../utils/trpc";
import Script from "next/script";

const MyApp: AppType<{ session: Session | null }> = ({
  Component,
  pageProps: { session, ...pageProps },
}) => {
  return (
    <>
      <Script
        strategy="afterInteractive"
        src={`https://www.googletagmanager.com/gtag/js?id=G-0N948SR48C`}
      />
      <Script id="google-analytics" strategy="afterInteractive">
        {`
          window.dataLayer = window.dataLayer || [];
          function gtag(){window.dataLayer.push(arguments);}
          gtag('js', new Date());

          gtag('config', 'G-0N948SR48C');
      `}
      </Script>
      <SessionProvider
        session={session}
        refetchInterval={60 * 60}
        refetchOnWindowFocus={false}
      >
        <Component {...pageProps} />
      </SessionProvider>
    </>
  );
};

// Update tRPC to v10: config() moved into createTRPCNext in utils/trpc
// and withTRPC<AppRouter>({...}) has been replaced by withTRPC(MyApp)
// Removed the withTRPC import above, used trpc.*
export default trpc.withTRPC(MyApp);
