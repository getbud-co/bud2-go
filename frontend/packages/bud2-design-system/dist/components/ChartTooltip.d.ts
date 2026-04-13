interface ChartTooltipEntry {
    name?: string;
    value?: number;
    color?: string;
    dataKey?: string | number;
}
interface ChartTooltipContentProps {
    active?: boolean;
    payload?: ChartTooltipEntry[];
    label?: string | number;
    labelFormatter?: (label: string) => string;
    valueFormatter?: (value: number) => string;
}
export declare function ChartTooltipContent({ active, payload, label, labelFormatter, valueFormatter, }: ChartTooltipContentProps): import("react/jsx-runtime").JSX.Element | null;
export {};
