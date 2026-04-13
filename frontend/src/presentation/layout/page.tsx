"use client";

import { useState, useMemo, useEffect, type ReactNode } from "react";
import { usePathname } from "next/navigation";

import { SidebarContext } from "@/contexts/SidebarContext";
import { SavedViewsProvider } from "@/contexts/SavedViewsContext";
import { AppSidebar } from "./layout-sidebar";

function useIsMobile(breakpoint = 768) {
  const [isMobile, setIsMobile] = useState(
    () => typeof window !== "undefined" && window.innerWidth <= breakpoint,
  );

  useEffect(() => {
    const mq = window.matchMedia(`(max-width: ${breakpoint}px)`);
    const handler = (e: MediaQueryListEvent) => setIsMobile(e.matches);
    mq.addEventListener("change", handler);
    setIsMobile(mq.matches);
    return () => mq.removeEventListener("change", handler);
  }, [breakpoint]);

  return isMobile;
}

interface AppLayoutProps {
  children: ReactNode;
}

export function AppLayout({ children }: AppLayoutProps) {
  const pathname = usePathname();
  const isMobile = useIsMobile();
  const [collapsed, setCollapsed] = useState(false);
  const [sidebarOpen, setSidebarOpen] = useState(false);

  const sidebarCtx = useMemo(
    () => ({ isMobile, openSidebar: () => setSidebarOpen(true) }),
    [isMobile],
  );

  if (pathname === "/login") {
    return <main className="min-h-dvh">{children}</main>;
  }

  return (
    <SavedViewsProvider>
      <SidebarContext.Provider value={sidebarCtx}>
        <div
          className="flex h-dvh bg-[var(--color-caramel-50)]"
          data-sidebar-collapsed={(!isMobile && collapsed) || undefined}
        >
          <AppSidebar
            collapsed={isMobile ? false : collapsed}
            onToggleCollapse={() => setCollapsed((prev) => !prev)}
            mobileOpen={isMobile ? sidebarOpen : undefined}
            onMobileClose={isMobile ? () => setSidebarOpen(false) : undefined}
          />
          <main
            className="flex-1 min-w-0 p-2! overflow-y-auto bg-[var(--color-caramel-50)] transition-[flex] duration-300 ease-in-out"
            data-scroll-region
          >
            {children}
          </main>
        </div>
      </SidebarContext.Provider>
    </SavedViewsProvider>
  );
}
