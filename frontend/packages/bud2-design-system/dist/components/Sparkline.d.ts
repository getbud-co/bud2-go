type SparklineColor = "orange" | "green" | "red" | "neutral";
interface SparklineProps {
    data: number[];
    width?: number;
    height?: number;
    color?: SparklineColor;
    filled?: boolean;
    className?: string;
}
export declare function Sparkline({ data, width, height, color, filled, className, }: SparklineProps): import("react/jsx-runtime").JSX.Element | null;
export {};
