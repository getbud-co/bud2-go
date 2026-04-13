export interface HeatmapCell {
    /** Row index or label */
    row: string;
    /** Column index or label */
    col: string;
    /** Numeric value for color intensity */
    value: number;
}
export interface HeatmapProps {
    /** Array of cells with row, col, and value */
    data: HeatmapCell[];
    /** Row labels in display order */
    rows: string[];
    /** Column labels in display order */
    columns: string[];
    /** Minimum value for color scale (default: auto from data) */
    min?: number;
    /** Maximum value for color scale (default: auto from data) */
    max?: number;
    /** Color palette (default: "orange"). Ignored when colorScale is "divergent". */
    color?: "orange" | "green" | "red" | "yellow" | "wine" | "neutral";
    /** Color scale mode (default: "sequential"). "divergent" uses red→yellow→green for low→mid→high values. */
    colorScale?: "sequential" | "divergent";
    /** Show values inside cells (default: true) */
    showValues?: boolean;
    /** Format function for cell values */
    formatValue?: (value: number) => string;
    /** Cell size in px (default: 40) */
    cellSize?: number;
    /** Row label column width in px (default: 100) */
    labelWidth?: number;
    /** Gap between cells in px (default: 4) */
    gap?: number;
    /** Map column labels to full-text tooltips shown on hover */
    columnTooltips?: Record<string, string>;
    className?: string;
}
export declare function Heatmap({ data, rows, columns, min: minProp, max: maxProp, color, colorScale, showValues, formatValue, cellSize, labelWidth, gap, columnTooltips, className, }: HeatmapProps): import("react/jsx-runtime").JSX.Element;
