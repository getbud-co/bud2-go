type ChartVariant = "full" | "half";
interface ChartProps {
    value: number;
    variant?: ChartVariant;
    /** Diameter in px (default 40). Controls the overall size of the gauge. */
    size?: number;
    className?: string;
}
export declare function Chart({ value, variant, size, className }: ChartProps): import("react/jsx-runtime").JSX.Element;
export {};
