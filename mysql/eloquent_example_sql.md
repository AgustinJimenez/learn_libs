select 
	* 
from 
	`ope04__remanencias` 
where 
	(
		select 
			count(*) 
		from 
			`ope04__cabeceras` 
		where 
			`ope04__remanencias`.`cabecera_codigo_sync` = `ope04__cabeceras`.`codigo_sync` 
		and 
		(
            select 
				count(*) 
			from 
				`formularios` 
			where 
				`ope04__cabeceras`.`codigo_formulario_sync` = `formularios`.`codigo_sync` 
			and 
			(
				select 
					count(*) 
				from 
					`formulario__travesia` 
				where 
					`formularios`.`formulario_travesia_id` = `formulario__travesia`.`id` 
				and 
				(
					select 
						count(*) 
					from 
						`travesia__travesias` 
					where 
						`formulario__travesia`.`travesia_id` = `travesia__travesias`.`id` 
					and 
					(
						select 
							count(*) 
						from 
							`unidad__unidads` 
						where 
							`travesia__travesias`.`unidad_id` = `unidad__unidads`.`id` 
						and 
                        (
							select 
								count(*) 
							from 
								`travesia__travesias` 
							where 
								`travesia__travesias`.`unidad_id` = `unidad__unidads`.`id` 
							and 
								`id` = 115
						) >= 1
					) >= 1
				) >= 1
			) >= 1
		) >= 1
	) >= 1 
order by 
	`formularios`.`codigo_sync` desc, `fecha` desc, `codigo_sync` desc
                            