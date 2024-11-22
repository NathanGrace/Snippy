<script lang="ts">
	import { Calendar } from 'bits-ui';
</script>

<Calendar.Root
	class="flex-col p-5 border rounded-2xl mx-auto max-w-sm "
	let:months
	let:weekdays
	weekdayFormat="long">
	<Calendar.Header class="flex items-center justify-between">
		<Calendar.PrevButton class="inline-flex size-10 items-center justify-center">
			<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
					 class="size-6">
				<path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5 8.25 12l7.5-7.5" />
			</svg>
		</Calendar.PrevButton>
		<Calendar.Heading class="text-base font-medium" />
		<Calendar.NextButton class="inline-flex size-10 items-center justify-center">
			<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
					 class="size-6">
				<path stroke-linecap="round" stroke-linejoin="round" d="m8.25 4.5 7.5 7.5-7.5 7.5" />
			</svg>
		</Calendar.NextButton>
	</Calendar.Header>
	<div class="pt-4">
		{#each months as month, i (i)}
			<Calendar.Grid class="w-full">
				<Calendar.GridHead>
					<Calendar.GridRow class="flex w-full">
						{#each weekdays as day}
							<Calendar.HeadCell class="w-full text-md">
								<div>{day.slice(0, 2)}</div>
							</Calendar.HeadCell>
						{/each}
					</Calendar.GridRow>
				</Calendar.GridHead>
				<Calendar.GridBody>
					{#each month.weeks as weekDates}
						<Calendar.GridRow class="flex w-full justify-between">
							{#each weekDates as date}
								<Calendar.Cell
									{date}
									class="!p-0"
								>
									<Calendar.Day
										{date}
										month={month.value}
										class="inline-flex size-12 items-center justify-center"
									>
									</Calendar.Day>
								</Calendar.Cell>
							{/each}
						</Calendar.GridRow>
					{/each}
				</Calendar.GridBody>
			</Calendar.Grid>
		{/each}
	</div>
</Calendar.Root>



<div class="p-5 border mt-4 mx-auto max-w-2xl overflow-y-auto rounded-2xl">
	<table class="table">
		<!-- head -->
		<thead>
		<tr>
			<th colspan="4">
				<button class="btn">Add New Appointment</button>
			</th>
		</tr>
		<tr>
			<th></th>
			<th>Time</th>
			<th>Name</th>
			<th>Service</th>
		</tr>
		</thead>
		
		<tbody>
<!--		{#each appointmentList as row, i}-->
		<!-- row 1 -->
		<tr class="hover">
			<th>
				<label>
					<button class="btn btn-xs btn-accent">Edit</button>
				</label>
			</th>
			<td>1:00pm</td>
			<td>Quality Control Specialist</td>
			<td>Blue</td>
		</tr>
		<!-- row 2 -->
		<tr class="hover">
			<th>
				<button class="btn btn-xs btn-accent">Edit</button>
			</th>
			<td>2:00pm</td>
			<td>Desktop Support Technician</td>
			<td>Purple</td>
		</tr>
		<!-- row 3 -->
		<tr class="hover">
			<th>
				<label>
					<button class="btn btn-xs btn-accent">Edit</button>
				</label>
			</th>
			<td>3:00pm</td>
			<td>Tax Accountant</td>
			<td>Red</td>
		</tr>
<!--		{/each}-->
		</tbody>
	</table>
</div>
