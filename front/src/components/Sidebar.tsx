import Image from "next/image";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
} from "./ui/sidebar";
import {
  Home,
  Calendar,
  LineChart,
  Landmark,
  Calculator,
  BanknoteArrowDown,
  BanknoteArrowUp,
  ChartCandlestick,
  Receipt,
  ChevronRight,
} from "lucide-react";
import { usePathname } from "next/navigation";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "./ui/collapsible";

const appVersion = process.env.APP_VERSION;
const items = [
  {
    title: "Resumo",
    url: "/",
    icon: Home,
  },
  {
    title: "Movimentações",
    url: "#",
    icon: Receipt,
    itens: [
      {
        title: "Entradas",
        url: "/revenue",
        icon: BanknoteArrowUp,
      },
      {
        title: "Saídas",
        url: "/expense",
        icon: BanknoteArrowDown,
      },
      {
        title: "Aporte",
        url: "/investments",
        icon: LineChart,
      },
    ],
  },
  {
    title: "Contas",
    url: "/mensal",
    icon: Calendar,
  },
  {
    title: "Previsão",
    url: "/previsao",
    icon: ChartCandlestick,
  },
  {
    title: "Patrimônio",
    url: "/patrimonio",
    icon: Landmark,
  },
  {
    title: "Simulador",
    url: "/simulador",
    icon: Calculator,
  },
];
function getMenuItens(item, path) {
  const subMenuItens = item.itens;

  if (subMenuItens && subMenuItens.length > 0) {
    const isActive = subMenuItens.some((sub) => sub.url === path);

    return (
      <Collapsible
        key={item.title}
        defaultOpen={isActive}
        className="group/collapsible"
      >
        <SidebarMenuItem>
          <CollapsibleTrigger asChild>
            <SidebarMenuButton>
              <item.icon />
              <span className="flex items-center gap-2 font-display text-xl">
                {item.title}
              </span>
              <ChevronRight className="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90" />
            </SidebarMenuButton>
          </CollapsibleTrigger>
          <CollapsibleContent>
            <SidebarMenuSub>
              {subMenuItens.map((subItem) => (
                <SidebarMenuSubItem key={subItem.title}>
                  <SidebarMenuSubButton asChild isActive={path === subItem.url}>
                    <a
                      href={subItem.url}
                      className="flex items-center gap-2 font-display text-xl"
                    >
                      <subItem.icon />
                      <span>{subItem.title}</span>
                    </a>
                  </SidebarMenuSubButton>
                </SidebarMenuSubItem>
              ))}
            </SidebarMenuSub>
          </CollapsibleContent>
        </SidebarMenuItem>
      </Collapsible>
    );
  }

  return (
    <SidebarMenuItem key={item.title}>
      <SidebarMenuButton asChild isActive={path === item.url}>
        <a
          href={item.url}
          className="flex items-center gap-2 font-display text-xl"
        >
          <item.icon />
          <span>{item.title}</span>
        </a>
      </SidebarMenuButton>
    </SidebarMenuItem>
  );
}
export function AppSidebar() {
  const pathname = usePathname();
  const menuGroups = [
    {
      label: "Organização",
      keys: ["Movimentações", "Contas"],
    },
    {
      label: "Construção",
      keys: ["Patrimônio"],
    },
    {
      label: "Futuro",
      keys: ["Previsão", "Simulador"],
    },
  ];
  const resumoItem = items.find((i) => i.title === "Resumo");
  return (
    <Sidebar>
      <SidebarHeader className="mt-4">
        <div className="flex items-center gap-3 px-2">
          <div className="flex aspect-square size-16 shrink-0 items-center justify-center rounded-lg bg-sidebar-primary/10 p-1">
            <Image
              src="/gofre-icon.png"
              width={500}
              height={500}
              alt="Gofre Logo"
              className="object-contain"
            />
          </div>

          <div className="flex flex-col min-w-0">
            <div className="flex items-center gap-2">
              <span className="font-display font-bold text-sm md:text-base truncate leading-tight">
                Gofre
              </span>
              <span className="rounded-md bg-muted px-1.5 py-0.5 font-number text-[10px] text-muted-foreground">
                v{appVersion}
              </span>
            </div>
            <span className="font-display truncate text-xs text-muted-foreground">
              Gestão Financeira
            </span>
          </div>
        </div>
      </SidebarHeader>
      <SidebarContent>
        {/* Resumo isolado */}
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu>
              {resumoItem && getMenuItens(resumoItem, pathname)}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>

        {/* Grupos */}
        {menuGroups.map((group) => (
          <SidebarGroup key={group.label}>
            <SidebarGroupLabel>{group.label}</SidebarGroupLabel>

            <SidebarGroupContent>
              <SidebarMenu>
                {items
                  .filter((item) => group.keys.includes(item.title))
                  .map((item) => getMenuItens(item, pathname))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        ))}
      </SidebarContent>
      <SidebarFooter />
    </Sidebar>
  );
}
