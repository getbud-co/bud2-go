import { PopoverItem } from './Popover';
interface RowActionsPopoverProps {
    /** Lista de itens do popover (ações disponíveis) */
    items: PopoverItem[];
    /** Se o popover está aberto */
    open: boolean;
    /** Callback para alternar estado aberto/fechado */
    onToggle: () => void;
    /** Callback quando o popover deve fechar */
    onClose: () => void;
    /** Classe CSS adicional para o wrapper */
    className?: string;
    /** Label acessível do botão de ações */
    buttonAriaLabel?: string;
}
/**
 * Wrapper padronizado para ações de linha de tabela.
 *
 * Combina um botão "⋯" (DotsThreeVertical) com um Popover de ações.
 * Simplifica o padrão repetitivo de botão + popover em tabelas CRUD.
 *
 * @example
 * ```tsx
 * const [openRowId, setOpenRowId] = useState<string | null>(null);
 *
 * function getActions(rowId: string): PopoverItem[] {
 *   return [
 *     { id: "edit", label: "Editar", icon: PencilSimple, onClick: () => handleEdit(rowId) },
 *     { id: "delete", label: "Excluir", icon: Trash, onClick: () => handleDelete(rowId) },
 *   ];
 * }
 *
 * // Na célula da tabela:
 * <TableCell align="right">
 *   <RowActionsPopover
 *     items={getActions(row.id)}
 *     open={openRowId === row.id}
 *     onToggle={() => setOpenRowId(openRowId === row.id ? null : row.id)}
 *     onClose={() => setOpenRowId(null)}
 *   />
 * </TableCell>
 * ```
 */
export declare function RowActionsPopover({ items, open, onToggle, onClose, className, buttonAriaLabel, }: RowActionsPopoverProps): import("react/jsx-runtime").JSX.Element;
export default RowActionsPopover;
