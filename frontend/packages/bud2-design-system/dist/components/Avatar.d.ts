export type AvatarSize = "xs" | "sm" | "md" | "lg" | "xl" | "2xl";
interface AvatarProps {
    /** Tamanho do avatar. */
    size?: AvatarSize;
    /** URL da foto. Quando ausente, mostra initials ou placeholder. */
    src?: string;
    /** Iniciais (1-2 caracteres). Usadas quando não há src. */
    initials?: string;
    /** Texto alternativo para a imagem. */
    alt?: string;
    /** Mostra indicador de online (bolinha verde). */
    online?: boolean;
    /** URL do logo da empresa (badge no canto inferior direito). */
    companyLogo?: string;
    className?: string;
}
export declare function Avatar({ size, src, initials, alt, online, companyLogo, className, }: AvatarProps): import("react/jsx-runtime").JSX.Element;
export {};
