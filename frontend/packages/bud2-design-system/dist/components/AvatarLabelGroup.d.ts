export type AvatarLabelGroupSize = "sm" | "md" | "lg" | "xl";
interface AvatarLabelGroupProps {
    /** Tamanho do grupo (sm | md | lg | xl). */
    size?: AvatarLabelGroupSize;
    /** Nome exibido ao lado do avatar. */
    name: string;
    /** Texto de apoio (e-mail, cargo, etc.). */
    supportingText?: string;
    /** URL da foto do avatar. */
    src?: string;
    /** Iniciais do avatar (fallback sem foto). */
    initials?: string;
    /** Texto alternativo da imagem. */
    alt?: string;
    /** Mostra indicador de online no avatar. */
    online?: boolean;
    /** URL do logo da empresa (badge no avatar). */
    companyLogo?: string;
    className?: string;
}
export declare function AvatarLabelGroup({ size, name, supportingText, src, initials, alt, online, companyLogo, className, }: AvatarLabelGroupProps): import("react/jsx-runtime").JSX.Element;
export {};
