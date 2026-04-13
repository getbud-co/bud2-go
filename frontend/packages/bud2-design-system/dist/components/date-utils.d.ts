export interface CalendarDate {
    year: number;
    /** 1-12 */
    month: number;
    /** 1-31 */
    day: number;
}
export declare const WEEKDAY_LABELS: string[];
export declare const MONTH_LABELS: string[];
export declare function daysInMonth(year: number, month: number): number;
/** 0 = Sunday … 6 = Saturday */
export declare function firstDayOfWeek(year: number, month: number): number;
export declare function prevMonth(d: CalendarDate): CalendarDate;
export declare function nextMonth(d: CalendarDate): CalendarDate;
export declare function isSameDay(a: CalendarDate, b: CalendarDate): boolean;
export declare function today(): CalendarDate;
export declare function isToday(d: CalendarDate): boolean;
/** -1 if a < b, 0 if equal, 1 if a > b */
export declare function compareDates(a: CalendarDate, b: CalendarDate): number;
export declare function isInRange(d: CalendarDate, start: CalendarDate, end: CalendarDate): boolean;
export declare function isDisabled(d: CalendarDate, minDate?: CalendarDate, maxDate?: CalendarDate): boolean;
export declare function formatDate(d: CalendarDate): string;
/** Parse DD/MM/AAAA → CalendarDate | null */
export declare function parseDate(str: string): CalendarDate | null;
export declare function isValidDate(year: number, month: number, day: number): boolean;
