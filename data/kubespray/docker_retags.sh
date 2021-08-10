names=$(docker images | grep -v "REPOSITORY" | awk '{print $1,$2}')
i=0
for name in ${names[@]}
do
	res=$(($i % 2))
	if [ $res == 0 ]; then
		old_tag=$name
		prefix=$(echo ${name#*/})
	else
		subfix=$(echo $name)
		new_tag=$(echo "harbor.dev.rdev.tech/kubesprary/$prefix:$subfix")
		old_tag=$old_tag:$subfix
		echo $old_tag
		echo $new_tag
		docker tag $old_tag $new_tag
		docker push $new_tag
	fi
	#if [ $res == 1 ]; then
	#	old_tag=''
	#	new_tag=''
	#	prefix=''
	#	subfix=''	
	#fi
	let i++
	
done
